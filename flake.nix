{
  description = "Quantum Cart Backend - Go Fiber Application";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix.url = "github:nix-community/gomod2nix";
    gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
    gomod2nix.inputs.utils.follows = "flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      gomod2nix,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ gomod2nix.overlays.default ];
        };

        # Primary application build
        app = pkgs.buildGoApplication {
          pname = "quantum-cart-backend";
          version = "0.1.0";
          src = ./.;
          modules = ./gomod2nix.toml;

          # Strip debug information for smaller binaries (like Dockerfile)
          ldflags = [
            "-s"
            "-w"
          ];

          # Environment variables required during build time if any
          CGO_ENABLED = 0;
        };

        # Docker image build using pkg's dockerTools
        dockerImage = pkgs.dockerTools.buildImage {
          name = "quantum-cart-backend";
          tag = "latest";

          # Use a minimal base like Alpine (or 'null' for scratch, but let's stick to alpine for shell access if needed)
          fromImage = pkgs.dockerTools.pullImage {
            imageName = "alpine";
            imageDigest = "sha256:51183f2cfa6320055da30872f211093f9ff1d3cf06f39a0bdb212"; # alpine:latest
            sha256 = "sha256-51183f2cfa6320055da30872f211093f9ff1d3cf06f39a0bdb212";
            finalImageTag = "latest";
            finalImageName = "alpine";
          };

          copyToRoot = pkgs.buildEnv {
            name = "image-root";
            paths = [
              app
              pkgs.cacert # Needed for Stripe/Twilio HTTPS calls
              pkgs.bash
              pkgs.coreutils
            ];
            pathsToLink = [
              "/bin"
              "/etc"
            ];
          };

          config = {
            Cmd = [ "${app}/bin/main" ];
            ExposedPorts = {
              "3000/tcp" = { };
            };
            Env = [
              "PORT=3000"
              "SSL_CERT_FILE=${pkgs.cacert}/etc/ssl/certs/ca-bundle.crt"
            ];
          };
        };

      in
      {
        packages.default = app;
        packages.docker = dockerImage;

        # Development environment
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
            go-tools
            gomod2nix.packages.${system}.default
            postgresql
          ];

          shellHook = ''
            echo "🚀 Quantum Cart Backend Dev Environment (Nix)"
            echo "Run 'gomod2nix' to update dependency lock file if you change go.mod"
          '';
        };
      }
    );
}

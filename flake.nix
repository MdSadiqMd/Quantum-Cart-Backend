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
          config.allowUnsupportedSystem = true;
        };

        # Primary application build
        app = pkgs.buildGoApplication {
          pname = "quantum-cart-backend";
          version = "0.1.0";
          src = ./.;
          modules = ./gomod2nix.toml;
          ldflags = [
            "-s"
            "-w"
          ];
          CGO_ENABLED = 0;
        };

        linux_app = app.overrideAttrs (old: {
          GOOS = "linux";
          GOARCH = "arm64";
        });

        pwd = builtins.getEnv "PWD";
        envFilePath = if pwd != "" then "${pwd}/.env" else ./.env;
        envFileResult = builtins.tryEval (builtins.readFile envFilePath);
        envFile = if envFileResult.success then envFileResult.value else "";

        # Create .env file in Nix store (will be copied to Docker image)
        # This keeps secrets out of the flake source code but includes them in the image
        envFileStore = pkgs.writeText "env-file" envFile;

        dockerImage = pkgs.dockerTools.buildLayeredImage {
          name = "quantum-cart-backend";
          tag = "latest";
          contents = [
            linux_app
            pkgs.cacert # Needed for Stripe/Twilio HTTPS calls
          ];

          extraCommands = ''
            mkdir -p app
            cp ${envFileStore} app/.env
            chmod 600 app/.env
          '';

          config = {
            Cmd = [ "${linux_app}/bin/linux_arm64/cmd" ];
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
        packages.linux_app = linux_app;
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

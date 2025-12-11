# Quantum Cart Backend

## 🛠️ Usage with Nix

This project is packaged with **Nix** for reproducible builds and zero-overhead deployments.

### 🧠 The Mental Model
*   **`flake.nix`**: The Blueprint. Defines *how* to build the app deterministically using Go inputs and nixpkgs.
*   **`gomod2nix.toml`**: The Translation. Maps Go dependencies (`go.sum`) to Nix-compatible hashes so builds are reproducible.

### 📝 How to Create & Use These Files
If you were to start from scratch, here is the flow:
1.  **Initialize Flake**: Create `flake.nix` (using the template from this repo)
    ```bash
    # Create a flake using the default template:
    nix flake init
    
    # List available templates:
    nix flake show templates

    # Create a flake from a specific template:
    nix flake init -t templates#simpleContainer
    ```
2.  **Generate Lock**: Nix needs to know your Go dependencies. Use `gomod2nix` to read your `go.mod` and write `gomod2nix.toml`:
    ```bash
    # Run this whenever go.mod changes
    nix run github:nix-community/gomod2nix -- generate
    ```

### 🚀 Running Locally & Developing
**Option 1: Build & Run Binary**
```bash
# Build the project (creates ./result)
nix build

# Run (load env vars first)
export $(cat .env | xargs) # Optional in most of the cases as env will already be loaded in the development shell
./result/bin/cmd
```

**Option 2: Development Shell**
Get a shell with Go, Postgres, and tools pre-installed (without polluting your system):
```bash
nix develop
# inside the shell:
go run cmd/main.go
```

### 🌐 Deploying / Running Elsewhere

#### **Linux (Zero Overhead)**
*   **With Nix**: `nix run github:MdSadiqMd/Quantum-Cart-Backend` (Automatically downloads & runs).
*   **Without Nix**: Copy the `./result/bin/cmd` binary to the server. It's statically linked and has **0 dependencies**

#### **Windows**
*   **WSL2**: Install Nix in WSL2 and follow Linux steps.
*   **Docker**:
    **Option A: Using Docker Compose (Recommended)**
    ```bash
    # Make sure you have a .env file with all required variables
    docker-compose up
    ```
    **Option B: Standalone Docker Run (Recommended)**
    ```bash
    # Build the Docker image (includes .env file from filesystem)
    # Note: Requires --impure flag to read untracked .env file
    nix build '.#docker' --impure
    docker load < result
    
    # Run the container (.env file is automatically loaded from /app/.env)
    docker run -p 3000:3000 quantum-cart-backend:latest
    ```
    **Option B2: Standalone Docker Run with External .env File**
    ```bash
    nix build '.#docker' --impure
    docker load < result
    
    # Run with environment variables from external .env file
    docker run -p 3000:3000 --env-file .env quantum-cart-backend:latest
    ```
    **Option C: Standalone Docker Run with Manual Env Vars**
    ```bash
    nix build .#docker
    docker load < result
    
    # Start PostgreSQL first (if using local database)
    docker run -d --name quantum_cart_db \
      -e POSTGRES_USER=postgres \
      -e POSTGRES_PASSWORD=postgres \
      -e POSTGRES_DB=quantum_cart \
      -p 5432:5432 \
      postgres:15-alpine
    
    # Run the app with all required environment variables
    docker run -p 3000:3000 \
      -e PORT=3000 \
      -e DB_URL=postgres://postgres:postgres@host.docker.internal:5432/quantum_cart?sslmode=disable \
      -e APP_SECRET=your_app_secret \
      -e TWILIO_SID=your_twilio_sid \
      -e TWILIO_AUTH_TOKEN=your_twilio_token \
      -e TWILIO_FROM_NUMBER=your_twilio_number \
      -e STRIPE_SECRET=your_stripe_secret \
      -e STRIPE_PUBLISHABLE_KEY=your_stripe_key \
      -e SUCCESS_URL=your_success_url \
      -e CANCEL_URL=your_cancel_url \
      quantum-cart-backend:latest
    ```
    
    *Note: This creates a "distroless" scratch image containing only the app and CA certs. It does not have `sh` or `bash` for debugging.*
    
    **For macOS/Windows**: Use `host.docker.internal` to connect to the host's PostgreSQL. For Linux, you may need to use `--network host` or the actual host IP.

### 🔄 Workflow Checklist
1.  **Modify Code/Deps**: Change `go.mod`? ->
2.  **Update Config**: `nix run github:nix-community/gomod2nix -- generate` ->
3.  **Build**: `nix build` (or `nix develop` to test)

### 🔐 Environment Variables & Secrets

The Docker image automatically includes your `.env` file during the build process:

- **Secrets are handled at the Nix level**: The `.env` file is read from your filesystem (requires `--impure` flag) and included in the Docker image at `/app/.env`
- **Secure by default**: The `.env` file is copied with secure permissions (600) and is not tracked in Git
- **Automatic loading**: Your application automatically loads environment variables from `/app/.env` at startup
- **No hardcoded secrets**: Secrets are never hardcoded in the flake source code

**Note**: The `--impure` flag is required when building the Docker image to allow Nix to read the untracked `.env` file from your filesystem.

### 📦 Changes & Fixes Log

**2. Nix-Level Secret Management:**
Secrets are now handled securely at the Nix build level:
*   **No Go-level encryption**: Removed encryption/decryption code from Go application
*   **Nix handles .env**: The `.env` file is read from filesystem and included in Docker image at build time
*   **Secure permissions**: `.env` file is copied with 600 permissions in the Docker image
*   **Requires --impure**: Building Docker image requires `--impure` flag to read untracked `.env` file

**1. "Exec Format Error" and Cross-Platform Fixes:**
Previously, we faced architectural mismatches (`exec format error`) when running the Docker image because we were mixing Linux/Arm64 binaries with Linux/Amd64 base images or running them on Darwin without proper cross-compilation handling.

**The Solution:**
*   **Removed Base Image (`alpine`)**: We now use a `scratch` (empty) base. This removes any dependency on upstream image hashes or architectures.
*   **Removed Hardcoded Architectures**: We removed `GOARCH=...` and `architecture=...`. Nix now automatically determines the correct architecture based on the building system.
*   **Removed System Binaries**: Access to `/bin/bash` or `/bin/ls` inside the container is removed. This prevents accidental copying of macOS binaries (Mach-O) into the Linux container.
*   **Result**: A pure, static, minimal (approx 15MB) Docker image that works on Linux and macOS (Arm64) natively.
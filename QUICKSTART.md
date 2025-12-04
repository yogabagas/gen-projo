# Quick Start - go-projo

## Installation (One-Time Setup)

```bash
cd go-projo
./install.sh
```

Or manually:
```bash
go install
```

## Basic Usage

```bash
go-projo gen -name <project-name> -module <module-path> -type <type>
```

## Examples

### Create a REST API
```bash
go-projo gen -name myapi -module github.com/user/myapi -type api
cd myapi
go mod tidy
make run
```

### Create a CLI Tool
```bash
go-projo gen -name mytool -module github.com/user/mytool -type cli
cd mytool
go mod tidy
make build
./bin/mytool
```

### Create a Microservice
```bash
go-projo gen -name myservice -module github.com/user/myservice -type microservice
cd myservice
go mod tidy
make docker-build
make docker-run
```

### Create a Library
```bash
go-projo gen -name mylib -module github.com/user/mylib -type library
cd mylib
go mod tidy
make test
```

## Commands

| Command | Description |
|---------|-------------|
| `go-projo gen` | Generate a new project |
| `go-projo version` | Show version |
| `go-projo help` | Show help |

## Gen Flags

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `-name` | ✅ Yes | - | Project name |
| `-module` | ✅ Yes | - | Go module path |
| `-type` | No | api | Project type (api, cli, microservice, library) |
| `-desc` | No | - | Project description |
| `-author` | No | - | Author name |
| `-go-version` | No | 1.24 | Go version |
| `-output` | No | . | Output directory |

## Project Types

| Type | Description | Use Case |
|------|-------------|----------|
| `api` | REST API server | Web APIs, backend services |
| `cli` | Command-line tool | CLI applications, scripts |
| `microservice` | Full microservice | Distributed systems, cloud-native apps |
| `library` | Go package | Reusable libraries, SDKs |

## After Generation

```bash
cd <project-name>
go mod tidy      # Install dependencies
make build       # Build the project
make test        # Run tests
make run         # Run the application
```

## Common Makefile Targets

| Command | Description |
|---------|-------------|
| `make build` | Build the application |
| `make run` | Run the application |
| `make test` | Run tests |
| `make coverage` | Show test coverage |
| `make clean` | Clean build artifacts |
| `make lint` | Run linter |

## API/Microservice Additional Targets

| Command | Description |
|---------|-------------|
| `make docker-build` | Build Docker image |
| `make docker-run` | Run in Docker |
| `make k8s-deploy` | Deploy to Kubernetes (microservice only) |

## Typical Workflow

1. **Generate project:**
   ```bash
   go-projo gen -name awesome-api -module github.com/me/awesome-api -type api
   ```

2. **Initialize:**
   ```bash
   cd awesome-api
   go mod tidy
   ```

3. **Develop:**
   - Add handlers in `internal/handler/`
   - Add business logic in `internal/service/`
   - Add data access in `internal/repository/`
   - Add models in `internal/model/`

4. **Test:**
   ```bash
   make test
   ```

5. **Run:**
   ```bash
   make run
   ```

6. **Deploy:**
   ```bash
   make docker-build
   make docker-run
   ```

## Project Structure (API)

```
myapi/
├── cmd/api/              # Entry point
├── internal/
│   ├── handler/          # HTTP handlers
│   ├── service/          # Business logic
│   ├── repository/       # Data access
│   ├── model/            # Data models
│   ├── middleware/       # Middleware
│   └── config/           # Config
├── pkg/
│   ├── response/         # Response helpers
│   └── validator/        # Validators
└── Makefile
```

## Tips

- **Module path:** Use your actual repository path (e.g., `github.com/username/project`)
- **Project type:** Choose based on your needs (api for web services, cli for tools)
- **Description:** Optional but helps document your project
- **Output:** Default is current directory, use `-output` to specify another location

## Troubleshooting

**go-projo command not found:**
```bash
# Check PATH
echo $PATH | grep $(go env GOPATH)/bin

# Add to PATH if missing
export PATH="$PATH:$(go env GOPATH)/bin"
```

**Build fails after generation:**
```bash
cd <project-name>
go mod tidy
go mod download
```

**Port already in use:**
- Edit `internal/config/config.go`
- Or set environment variable: `SERVER_ADDRESS=:3000`

## Help

```bash
go-projo help           # General help
go-projo gen -help      # Generate command help
```

## Full Documentation

- [README.md](README.md) - Complete documentation
- [INSTALL.md](INSTALL.md) - Installation guide

---

**Happy coding!** 🚀

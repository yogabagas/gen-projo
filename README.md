# go-projo - Go Project Structure Generator

A powerful CLI tool to generate well-structured Go projects following best practices.

## Installation

### Quick Install (Recommended)

```bash
cd go-projo
./install.sh
```

This will build and install `go-projo` to your `$GOPATH/bin` directory.

### Install via `go install`

```bash
go install github.com/yogabagas/gen-projo@latest
```

### Manual Install

```bash
cd go-projo
make install
```

**See [INSTALL.md](INSTALL.md) for detailed installation instructions and troubleshooting.**

## Usage

### Basic Command

```bash
go-projo gen -name <project-name> -module <module-path> [options]
```

### Commands

- `gen`, `generate` - Generate a new Go project
- `version` - Show version information
- `help` - Show help message

### Generate Command Options

**Required:**
- `-name` - Project name
- `-module` - Go module path (e.g., github.com/user/project)

**Optional:**
- `-type` - Project type: `api`, `cli`, `microservice`, `library` (default: "api")
- `-desc` - Project description
- `-author` - Author name
- `-go-version` - Go version (default: "1.24")
- `-output` - Output directory path (default: current directory)

## Quick Examples

### Generate REST API

```bash
go-projo gen -name myapi -module github.com/user/myapi -type api
```

### Generate CLI Tool

```bash
go-projo gen -name mytool -module github.com/user/mytool -type cli
```

### Generate Microservice

```bash
go-projo gen -name myservice -module github.com/user/myservice -type microservice
```

### Generate Library

```bash
go-projo gen -name mylib -module github.com/user/mylib -type library
```

### With All Options

```bash
go-projo gen \
  -name awesome-api \
  -module github.com/mycompany/awesome-api \
  -type api \
  -desc "My awesome REST API" \
  -author "John Doe" \
  -go-version 1.24 \
  -output ~/projects
```

## Project Types

### 1. API (REST API)
Creates a REST API project with:
- HTTP server with graceful shutdown
- Handler/Service/Repository layers
- Middleware (logging, CORS)
- Response utilities
- Health check endpoint
- Makefile with build tasks

**Structure:**
```
myapi/
├── cmd/api/              # Application entrypoint
├── internal/
│   ├── handler/          # HTTP handlers
│   ├── service/          # Business logic
│   ├── repository/       # Data access
│   ├── model/            # Domain models
│   ├── middleware/       # HTTP middleware
│   └── config/           # Configuration
├── pkg/
│   ├── response/         # Response helpers
│   └── validator/        # Validators
├── docs/                 # Documentation
├── Makefile
└── go.mod
```

### 2. CLI (Command Line Tool)
Creates a CLI application with:
- Command structure
- Configuration management
- Utility packages

**Structure:**
```
mytool/
├── cmd/                  # Main application
├── internal/
│   ├── command/          # CLI commands
│   └── config/           # Configuration
├── pkg/utils/            # Utilities
├── Makefile
└── go.mod
```

### 3. Microservice
Creates a microservice with:
- HTTP and gRPC server support
- Docker configuration
- Kubernetes manifests
- All features from API type

**Structure:**
```
myservice/
├── cmd/server/           # Server entrypoint
├── internal/             # Same as API
├── pkg/
│   ├── grpc/             # gRPC utilities
│   └── http/             # HTTP utilities
├── proto/                # Protocol buffers
├── deployments/
│   ├── docker/           # Docker configs
│   └── k8s/              # Kubernetes manifests
├── Dockerfile
├── Makefile
└── go.mod
```

### 4. Library
Creates a reusable Go library with:
- Clean package structure
- Example usage code
- Documentation

**Structure:**
```
mylib/
├── mylib.go              # Main library code
├── internal/             # Private code
├── examples/             # Usage examples
├── docs/                 # Documentation
├── Makefile
└── go.mod
```

## After Generation

Once your project is generated:

```bash
# Navigate to project
cd myproject

# Install dependencies
go mod tidy

# Build
make build

# Run (for API/Microservice/CLI)
make run

# Run tests
make test

# View coverage
make coverage
```

## Makefile Commands

All generated projects include these commands:

- `make build` - Build the application
- `make run` - Run the application
- `make test` - Run tests
- `make coverage` - View test coverage
- `make clean` - Clean build artifacts
- `make lint` - Run linter

API/Microservice projects also include:
- `make docker-build` - Build Docker image
- `make docker-run` - Run Docker container

Microservice projects also include:
- `make k8s-deploy` - Deploy to Kubernetes
- `make proto` - Generate protobuf code

## Features

✅ Multiple project types (API, CLI, Microservice, Library)
✅ Complete directory structure following Go best practices
✅ Ready-to-use boilerplate code
✅ Makefile with common tasks
✅ Docker and Kubernetes support
✅ Middleware and utilities
✅ Clean architecture (handler → service → repository)
✅ .gitignore and README included
✅ Interactive confirmation before generation

## Project Layout Philosophy

This tool follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):

- `cmd/` - Application entrypoints
- `internal/` - Private application code
- `pkg/` - Public libraries
- `docs/` - Documentation
- `deployments/` - Deployment configurations

## Examples

### Create and Run an API

```bash
# Generate
go-projo gen -name userapi -module github.com/company/userapi -type api

# Setup
cd userapi
go mod tidy

# Build and run
make build
make run

# Test health endpoint
curl http://localhost:8080/health
```

### Create and Run a CLI Tool

```bash
# Generate
go-projo gen -name datamigrate -module github.com/company/datamigrate -type cli

# Setup and build
cd datamigrate
go mod tidy
make build

# Run
./bin/datamigrate
```

### Create a Microservice with Docker

```bash
# Generate
go-projo gen -name authservice -module github.com/company/authservice -type microservice

# Setup
cd authservice
go mod tidy

# Run with Docker
make docker-build
make docker-run
```

## Version

Check version:
```bash
go-projo version
```

## Help

Show help:
```bash
go-projo help
go-projo gen -help
```

## Contributing

Feel free to submit issues and pull requests!

## License

MIT License

## Author

Created with ❤️ for the Go community

package generator

// Template strings for various file types

const goModTemplate = `module {{.Module}}

go {{.GoVersion}}

require (
	// Add your dependencies here
)
`

const readmeTemplate = `# {{.Name}}

{{.Description}}

## Author

{{.Author}}

## Getting Started

### Prerequisites

- Go {{.GoVersion}} or higher

### Installation

` + "```bash" + `
go get {{.Module}}
` + "```" + `

### Usage

` + "```bash" + `
# Build the project
make build

# Run tests
make test

# Run the application
make run
` + "```" + `

## Project Structure

` + "```" + `
{{.Name}}/
├── cmd/          # Application entrypoints
├── internal/     # Private application code
├── pkg/          # Public libraries
└── docs/         # Documentation
` + "```" + `

## License

MIT License
`

const gitignoreTemplate = `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with ` + "`go test -c`" + `
*.test

# Output of the go coverage tool
*.out

# Go workspace file
go.work

# Dependency directories
vendor/

# IDEs
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Build artifacts
bin/
dist/
build/

# Environment variables
.env
.env.local
.env.*.local

# Logs
*.log

# Temporary files
tmp/
temp/
`

const makefileAPITemplate = `.PHONY: build run test clean docker-build docker-run

APP_NAME={{.Name}}
VERSION?=latest
DOCKER_IMAGE={{.Name}}:${VERSION}

build:
	go build -o bin/${APP_NAME} cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v -race -coverprofile=coverage.out ./...

coverage:
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -f coverage.out

lint:
	golangci-lint run

docker-build:
	docker build -t ${DOCKER_IMAGE} .

docker-run:
	docker run -p 8080:8080 ${DOCKER_IMAGE}

migrate-up:
	# Add your migration command here

migrate-down:
	# Add your migration command here

help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  coverage     - Show test coverage"
	@echo "  clean        - Clean build artifacts"
	@echo "  lint         - Run linter"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
`

const makefileCLITemplate = `.PHONY: build install test clean

APP_NAME={{.Name}}

build:
	go build -o bin/${APP_NAME} cmd/main.go

install:
	go install cmd/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

lint:
	golangci-lint run
`

const makefileMicroTemplate = `.PHONY: build run test proto docker-build docker-run k8s-deploy

APP_NAME={{.Name}}
VERSION?=latest
DOCKER_IMAGE={{.Name}}:${VERSION}

build:
	go build -o bin/${APP_NAME} cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test -v -race ./...

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

docker-build:
	docker build -t ${DOCKER_IMAGE} .

docker-run:
	docker run -p 8080:8080 -p 9090:9090 ${DOCKER_IMAGE}

k8s-deploy:
	kubectl apply -f deployments/k8s/

k8s-delete:
	kubectl delete -f deployments/k8s/

clean:
	rm -rf bin/
`

const makefileLibTemplate = `.PHONY: test coverage lint example

test:
	go test -v -race ./...

coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	golangci-lint run

example:
	go run examples/main.go
`

const mainAPITemplate = `package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{.Module}}/internal/config"
	"{{.Module}}/internal/handler"
	"{{.Module}}/internal/middleware"
	"{{.Module}}/internal/repository"
	"{{.Module}}/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize repository
	repo := repository.New()

	// Initialize service
	svc := service.New(repo)

	// Initialize handler
	h := handler.New(svc)

	// Setup router
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/api/v1/", h.HandleAPI)

	// Apply middleware
	handler := middleware.Logger(mux)
	handler = middleware.CORS(handler)

	// Create server
	srv := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server starting on %s", cfg.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
`

const mainCLITemplate = `package main

import (
	"fmt"
	"os"

	"{{.Module}}/internal/command"
)

func main() {
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
`

const mainMicroTemplate = `package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"{{.Module}}/internal/config"
	"{{.Module}}/internal/handler"
	"{{.Module}}/internal/repository"
	"{{.Module}}/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize layers
	repo := repository.New()
	svc := service.New(repo)
	h := handler.New(svc)

	// Start HTTP server
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/health", h.Health)
		log.Printf("HTTP server listening on %s", cfg.HTTPAddress)
		if err := http.ListenAndServe(cfg.HTTPAddress, mux); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Start gRPC server (if needed)
	go func() {
		lis, err := net.Listen("tcp", cfg.GRPCAddress)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on %s", cfg.GRPCAddress)
		// Initialize gRPC server here
		_ = lis
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")
	// Add cleanup logic
}
`

const configTemplate = `package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	HTTPAddress   string
	GRPCAddress   string
	Environment   string
}

func Load() (*Config, error) {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		HTTPAddress:   getEnv("HTTP_ADDRESS", ":8080"),
		GRPCAddress:   getEnv("GRPC_ADDRESS", ":9090"),
		Environment:   getEnv("ENVIRONMENT", "development"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
`

const handlerTemplate = `package handler

import (
	"encoding/json"
	"net/http"

	"{{.Module}}/internal/service"
	"{{.Module}}/pkg/response"
)

type Handler struct {
	service *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

func (h *Handler) HandleAPI(w http.ResponseWriter, r *http.Request) {
	// Implement your API handlers here
	response.JSON(w, http.StatusOK, map[string]string{
		"message": "API endpoint",
	})
}
`

const serviceTemplate = `package service

import (
	"{{.Module}}/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Add your business logic methods here
`

const repositoryTemplate = `package repository

type Repository struct {
	// Add your database connections here
}

func New() *Repository {
	return &Repository{}
}

// Add your data access methods here
`

const modelTemplate = `package model

// Add your domain models here

type Example struct {
	ID   string ` + "`json:\"id\"`" + `
	Name string ` + "`json:\"name\"`" + `
}
`

const middlewareTemplate = `package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
`

const responseTemplate = `package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        ` + "`json:\"success\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
	Error   string      ` + "`json:\"error,omitempty\"`" + `
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Success: status < 400,
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}

func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Success: false,
		Error:   message,
	}

	json.NewEncoder(w).Encode(resp)
}
`

const cliRootTemplate = `package command

import (
	"fmt"
)

func Execute() error {
	// Implement your CLI commands here
	fmt.Println("{{.Name}} CLI")
	return nil
}
`

const apiDocsTemplate = `# API Documentation

## Endpoints

### Health Check

` + "```" + `
GET /health
` + "```" + `

Returns the health status of the service.

### API v1

` + "```" + `
GET /api/v1/
` + "```" + `

Main API endpoint.
`

const dockerfileTemplate = `FROM golang:{{.GoVersion}}-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /{{.Name}} cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /{{.Name}} .

EXPOSE 8080 9090

CMD ["./{{.Name}}"]
`

const k8sDeploymentTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
spec:
  replicas: 3
  selector:
    matchLabels:
      app: {{.Name}}
  template:
    metadata:
      labels:
        app: {{.Name}}
    spec:
      containers:
      - name: {{.Name}}
        image: {{.Name}}:latest
        ports:
        - containerPort: 8080
        - containerPort: 9090
        env:
        - name: ENVIRONMENT
          value: "production"
`

const k8sServiceTemplate = `apiVersion: v1
kind: Service
metadata:
  name: {{.Name}}
spec:
  selector:
    app: {{.Name}}
  ports:
  - name: http
    port: 80
    targetPort: 8080
  - name: grpc
    port: 9090
    targetPort: 9090
  type: LoadBalancer
`

const libraryMainTemplate = `package {{.Name}}

// Add your library implementation here

type Client struct {
	// Configuration fields
}

func New() *Client {
	return &Client{}
}
`

const libraryExampleTemplate = `package main

import (
	"fmt"

	"{{.Module}}"
)

func main() {
	client := {{.Name}}.New()
	fmt.Printf("{{.Name}} client: %+v\n", client)
}
`

const usageDocsTemplate = `# Usage Guide

## Installation

` + "```bash" + `
go get {{.Module}}
` + "```" + `

## Basic Usage

` + "```go" + `
package main

import "{{.Module}}"

func main() {
    client := {{.Name}}.New()
    // Use the client
}
` + "```" + `

## Examples

See the [examples](../examples/) directory for more usage examples.
`

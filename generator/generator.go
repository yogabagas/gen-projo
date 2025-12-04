package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ProjectType represents different types of Go projects
type ProjectType string

const (
	ProjectTypeAPI     ProjectType = "api"
	ProjectTypeCLI     ProjectType = "cli"
	ProjectTypeMicro   ProjectType = "microservice"
	ProjectTypeLibrary ProjectType = "library"
)

// ProjectConfig holds configuration for project generation
type ProjectConfig struct {
	Name        string
	Module      string
	Type        ProjectType
	Description string
	Author      string
	GoVersion   string
	OutputPath  string
}

// ProjectStructure defines the directory and file structure
type ProjectStructure struct {
	Directories []string
	Files       map[string]string
}

// Generator handles project structure generation
type Generator struct {
	config    ProjectConfig
	structure ProjectStructure
}

// NewGenerator creates a new Generator instance
func NewGenerator(config ProjectConfig) *Generator {
	if config.GoVersion == "" {
		config.GoVersion = "1.24"
	}

	g := &Generator{
		config: config,
	}

	g.buildStructure()
	return g
}

// buildStructure builds the project structure based on project type
func (g *Generator) buildStructure() {
	switch g.config.Type {
	case ProjectTypeAPI:
		g.structure = g.buildAPIStructure()
	case ProjectTypeCLI:
		g.structure = g.buildCLIStructure()
	case ProjectTypeMicro:
		g.structure = g.buildMicroserviceStructure()
	case ProjectTypeLibrary:
		g.structure = g.buildLibraryStructure()
	default:
		g.structure = g.buildAPIStructure()
	}
}

// Generate creates the project structure on disk
func (g *Generator) Generate() error {
	basePath := filepath.Join(g.config.OutputPath, g.config.Name)

	// Create base directory
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return fmt.Errorf("failed to create base directory: %w", err)
	}

	// Create all directories
	for _, dir := range g.structure.Directories {
		dirPath := filepath.Join(basePath, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create all files
	for filePath, content := range g.structure.Files {
		fullPath := filepath.Join(basePath, filePath)

		// Parse template
		tmpl, err := template.New(filePath).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse template for %s: %w", filePath, err)
		}

		// Create file
		file, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
		defer file.Close()

		// Execute template
		if err := tmpl.Execute(file, g.config); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
	}

	return nil
}

// buildAPIStructure creates structure for REST API projects
func (g *Generator) buildAPIStructure() ProjectStructure {
	return ProjectStructure{
		Directories: []string{
			"cmd/api",
			"internal/handler",
			"internal/service",
			"internal/repository",
			"internal/model",
			"internal/middleware",
			"internal/config",
			"pkg/response",
			"pkg/validator",
			"migrations",
			"docs",
			"scripts",
		},
		Files: map[string]string{
			"go.mod":                        goModTemplate,
			"README.md":                     readmeTemplate,
			".gitignore":                    gitignoreTemplate,
			"Makefile":                      makefileAPITemplate,
			"cmd/api/main.go":               mainAPITemplate,
			"internal/config/config.go":     configTemplate,
			"internal/handler/handler.go":   handlerTemplate,
			"internal/service/service.go":   serviceTemplate,
			"internal/repository/repository.go": repositoryTemplate,
			"internal/model/model.go":       modelTemplate,
			"internal/middleware/middleware.go": middlewareTemplate,
			"pkg/response/response.go":      responseTemplate,
			"docs/API.md":                   apiDocsTemplate,
		},
	}
}

// buildCLIStructure creates structure for CLI tool projects
func (g *Generator) buildCLIStructure() ProjectStructure {
	return ProjectStructure{
		Directories: []string{
			"cmd",
			"internal/command",
			"internal/config",
			"pkg/utils",
			"docs",
		},
		Files: map[string]string{
			"go.mod":                    goModTemplate,
			"README.md":                 readmeTemplate,
			".gitignore":                gitignoreTemplate,
			"Makefile":                  makefileCLITemplate,
			"cmd/main.go":               mainCLITemplate,
			"internal/command/root.go":  cliRootTemplate,
			"internal/config/config.go": configTemplate,
		},
	}
}

// buildMicroserviceStructure creates structure for microservice projects
func (g *Generator) buildMicroserviceStructure() ProjectStructure {
	return ProjectStructure{
		Directories: []string{
			"cmd/server",
			"internal/handler",
			"internal/service",
			"internal/repository",
			"internal/model",
			"internal/middleware",
			"internal/config",
			"pkg/grpc",
			"pkg/http",
			"proto",
			"migrations",
			"deployments/docker",
			"deployments/k8s",
			"scripts",
		},
		Files: map[string]string{
			"go.mod":                        goModTemplate,
			"README.md":                     readmeTemplate,
			".gitignore":                    gitignoreTemplate,
			"Makefile":                      makefileMicroTemplate,
			"Dockerfile":                    dockerfileTemplate,
			"cmd/server/main.go":            mainMicroTemplate,
			"internal/config/config.go":     configTemplate,
			"internal/handler/handler.go":   handlerTemplate,
			"internal/service/service.go":   serviceTemplate,
			"internal/repository/repository.go": repositoryTemplate,
			"internal/model/model.go":       modelTemplate,
			"deployments/k8s/deployment.yaml":   k8sDeploymentTemplate,
			"deployments/k8s/service.yaml":      k8sServiceTemplate,
		},
	}
}

// buildLibraryStructure creates structure for library projects
func (g *Generator) buildLibraryStructure() ProjectStructure {
	return ProjectStructure{
		Directories: []string{
			"internal",
			"examples",
			"docs",
		},
		Files: map[string]string{
			"go.mod":          goModTemplate,
			"README.md":       readmeTemplate,
			".gitignore":      gitignoreTemplate,
			"Makefile":        makefileLibTemplate,
			"{{.Name}}.go":    libraryMainTemplate,
			"examples/main.go": libraryExampleTemplate,
			"docs/USAGE.md":   usageDocsTemplate,
		},
	}
}

// GetProjectInfo returns formatted project information
func (g *Generator) GetProjectInfo() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Project: %s\n", g.config.Name))
	sb.WriteString(fmt.Sprintf("Module: %s\n", g.config.Module))
	sb.WriteString(fmt.Sprintf("Type: %s\n", g.config.Type))
	sb.WriteString(fmt.Sprintf("Go Version: %s\n", g.config.GoVersion))
	sb.WriteString(fmt.Sprintf("Output Path: %s\n", filepath.Join(g.config.OutputPath, g.config.Name)))
	sb.WriteString(fmt.Sprintf("\nDirectories: %d\n", len(g.structure.Directories)))
	sb.WriteString(fmt.Sprintf("Files: %d\n", len(g.structure.Files)))

	return sb.String()
}

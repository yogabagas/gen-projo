package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yourusername/go-projo/generator"
)

func executeGenerate() error {
	// Create a new FlagSet for the generate command
	fs := flag.NewFlagSet("gen", flag.ExitOnError)

	var (
		name        = fs.String("name", "", "Project name (required)")
		module      = fs.String("module", "", "Go module path (required)")
		projectType = fs.String("type", "api", "Project type: api, cli, microservice, library")
		description = fs.String("desc", "", "Project description")
		author      = fs.String("author", "", "Author name")
		goVersion   = fs.String("go-version", "1.24", "Go version")
		outputPath  = fs.String("output", ".", "Output directory path")
		help        = fs.Bool("help", false, "Show help message")
	)

	// Custom usage function
	fs.Usage = func() {
		showGenerateHelp()
	}

	// Parse flags starting from os.Args[2] (after 'gen')
	if err := fs.Parse(os.Args[2:]); err != nil {
		return err
	}

	// Show help if requested
	if *help {
		showGenerateHelp()
		return nil
	}

	// Validate required flags
	if *name == "" {
		return fmt.Errorf("-name is required\n\nRun 'go-projo gen -help' for usage")
	}

	if *module == "" {
		return fmt.Errorf("-module is required\n\nRun 'go-projo gen -help' for usage")
	}

	// Validate project type
	var pType generator.ProjectType
	switch *projectType {
	case "api":
		pType = generator.ProjectTypeAPI
	case "cli":
		pType = generator.ProjectTypeCLI
	case "microservice", "micro":
		pType = generator.ProjectTypeMicro
	case "library", "lib":
		pType = generator.ProjectTypeLibrary
	default:
		return fmt.Errorf("invalid project type '%s'. Must be one of: api, cli, microservice, library", *projectType)
	}

	// Get absolute output path
	absOutputPath, err := filepath.Abs(*outputPath)
	if err != nil {
		return fmt.Errorf("invalid output path: %v", err)
	}

	// Create generator config
	config := generator.ProjectConfig{
		Name:        *name,
		Module:      *module,
		Type:        pType,
		Description: *description,
		Author:      *author,
		GoVersion:   *goVersion,
		OutputPath:  absOutputPath,
	}

	// Create generator
	gen := generator.NewGenerator(config)

	// Show project info
	fmt.Println("=== Go Project Generator ===")
	fmt.Println()
	fmt.Print(gen.GetProjectInfo())
	fmt.Println()

	// Confirm generation
	fmt.Print("Generate project? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("Generation cancelled")
		return nil
	}

	// Generate project
	fmt.Println("\nGenerating project...")
	if err := gen.Generate(); err != nil {
		return fmt.Errorf("failed to generate project: %v", err)
	}

	fmt.Println("✓ Project generated successfully!")
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("  cd %s\n", filepath.Join(absOutputPath, *name))
	fmt.Printf("  go mod tidy\n")
	fmt.Printf("  make build\n")

	return nil
}

func showGenerateHelp() {
	fmt.Println(`Generate a new Go project structure

Usage:
  go-projo gen [flags]

Flags:
  -name string
        Project name (required)
  -module string
        Go module path (required)
  -type string
        Project type: api, cli, microservice, library (default "api")
  -desc string
        Project description
  -author string
        Author name
  -go-version string
        Go version (default "1.24")
  -output string
        Output directory path (default ".")
  -help
        Show this help message

Project Types:
  api           REST API server with HTTP handlers
  cli           Command-line tool
  microservice  Microservice with HTTP/gRPC and Docker/K8s configs
  library       Reusable Go library package

Examples:
  # Generate REST API project
  go-projo gen -name myapi -module github.com/user/myapi -type api

  # Generate CLI tool project
  go-projo gen -name mytool -module github.com/user/mytool -type cli

  # Generate microservice project with description
  go-projo gen -name myservice -module github.com/user/myservice -type microservice -desc "My awesome service"

  # Generate library project
  go-projo gen -name mylib -module github.com/user/mylib -type library -author "Your Name"

  # Generate to a specific directory
  go-projo gen -name myapi -module github.com/user/myapi -output ~/projects`)
}

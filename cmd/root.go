package cmd

import (
	"fmt"
	"os"
)

var version = "1.0.0"

// Execute runs the CLI application
func Execute() error {
	if len(os.Args) < 2 {
		showHelp()
		return nil
	}

	switch os.Args[1] {
	case "gen", "generate":
		return executeGenerate()
	case "version", "-v", "--version":
		fmt.Printf("go-projo version %s\n", version)
		return nil
	case "help", "-h", "--help":
		showHelp()
		return nil
	default:
		return fmt.Errorf("unknown command: %s\nRun 'go-projo help' for usage", os.Args[1])
	}
}

func showHelp() {
	fmt.Println(`go-projo - Go Project Structure Generator

Usage:
  go-projo <command> [flags]

Commands:
  gen, generate    Generate a new Go project
  version          Show version information
  help             Show this help message

Examples:
  go-projo gen -name myapi -module github.com/user/myapi -type api
  go-projo gen -name mytool -module github.com/user/mytool -type cli
  go-projo version

Run 'go-projo gen -help' for more information about the generate command.`)
}

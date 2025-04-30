package main

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

var SonarToken string

func init() {
	SonarToken = os.Getenv("SONAR_TOKEN")
	if SonarToken == "" {
		fmt.Println("SONAR_TOKEN environment variable is not set")
		os.Exit(1)
	}
}

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Sonarcloud API",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithRecovery(),
		server.WithLogging(),
	)

	// add the individual tools
	AddSonarProjectsTool(s)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Sonar MCP server error: %v\n", err)
	}
}

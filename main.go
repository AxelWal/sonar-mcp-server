package main

import (
	"flag"
	"log"

	"github.com/lreimer/sonar-mcp-server/sonar"
	"github.com/mark3labs/mcp-go/server"
)

var version = "dev"

func main() {
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")
	flag.Parse()

	// Create a new MCP server
	s := server.NewMCPServer(
		"Sonarcloud API",
		version,
		server.WithResourceCapabilities(true, true),
		server.WithRecovery(),
		server.WithLogging(),
	)

	// add the individual tools
	sonar.AddProjectsTool(s)
	sonar.AddHotspotsTool(s)
	sonar.AddIssuesTool(s)
	sonar.AddDuplicationsTool(s)

	// Only check for "sse" since stdio is the default
	if transport == "sse" {
		sseServer := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8080"))
		log.Printf("Sonar MCP Server (SSE) listening on :8080")
		if err := sseServer.Start(":8080"); err != nil {
			log.Fatalf("Sonar MCP Server (SSE) error: %v", err)
		}
	} else {
		if err := server.ServeStdio(s); err != nil {
			log.Fatalf("Sonar MCP Server (stdio) error: %v", err)
		}
	}
}

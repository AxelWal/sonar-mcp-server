package main

import (
	"flag"
	"log"

	"github.com/lreimer/sonar-mcp-server/sonar"
	"github.com/mark3labs/mcp-go/server"
)

var version = "dev"

func main() {
	var transport, port, baseURL string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")
	flag.StringVar(&port, "p", "8080", "Port for SSE transport")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "Base URL for SSE transport")
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
		sseServer := server.NewSSEServer(s, server.WithBaseURL(baseURL))
		ssePort := "0.0.0.0:" + port
		log.Printf("Sonar MCP Server (SSE) listening on :8080")
		if err := sseServer.Start(ssePort); err != nil {
			log.Fatalf("Sonar MCP Server (SSE) error: %v", err)
		}
	} else {
		if err := server.ServeStdio(s); err != nil {
			log.Fatalf("Sonar MCP Server (stdio) error: %v", err)
		}
	}
}

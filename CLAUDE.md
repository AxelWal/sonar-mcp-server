# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

Sonar MCP Server is a Go implementation of an MCP (Model Context Protocol) server for integrating with the SonarQube Cloud API. It's designed to be used with Claude and other AI assistants to provide information about SonarQube projects, issues, hotspots, and code duplications.

## Commands

### Build and Run

```bash
# Basic build
make build
# or
go build -ldflags "-X main.version=$(git describe --tags --always --dirty)"

# Clean build artifacts
make clean

# Quick snapshot build with goreleaser
goreleaser build --snapshot --clean

# Create a snapshot release with goreleaser
goreleaser release --skip=publish --snapshot --clean

# Run as stdio binary (default)
./sonar-mcp-server -t stdio

# Run as SSE server
./sonar-mcp-server -t sse -p 8080 -b http://localhost:8080
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./sonar

# Run a specific test
go test ./sonar -run TestUnmarshalProjectsResponse

# Run tests with verbose output
go test -v ./...
```

## Environment Setup

The server requires a SonarQube API token to be set as an environment variable:

```bash
export SONAR_TOKEN="your-sonar-token"
```

## Architecture

The codebase follows a simple modular structure:

1. **Main Package** (`main.go`): Entry point that sets up the MCP server with the appropriate transport (stdio or SSE) and registers all the Sonar API tools.

2. **Sonar Package** (`sonar/`): Contains all the functionality for interacting with the SonarQube API:
   - `types.go`: Defines data structures for SonarQube API responses
   - `projects.go`: Tool for listing SonarQube projects
   - `issues.go`: Tool for searching and retrieving project issues
   - `hotspots.go`: Tool for working with security hotspots
   - `duplications.go`: Tool for code duplication detection
   - `security.go`: Handles token security
   - `utils.go`: Common utility functions for HTTP requests and response formatting

3. **MCP Server Integration**: Uses the `mark3labs/mcp-go` package to implement the Model Context Protocol, allowing AI assistants to interact with the SonarQube API.

## Development Workflow

1. Each tool follows the same pattern:
   - Define data structures for API responses
   - Create a function to add the tool to the MCP server (`AddXxxTool`)
   - Implement the API call function (`searchXxx`)
   - Define parameters and their validation rules

2. HTTP requests are performed using the common `performGetRequest` function, which adds authentication using the SonarQube token.

3. The `getSonarToken` function in `security.go` retrieves the token from the environment.

## Deployment

The server can be deployed to Google Cloud Run using the commands provided in the README:

1. Create a secret for the SONAR_TOKEN
2. Enable required services
3. Set up IAM permissions for the secret
4. Deploy the server to Cloud Run

Alternatively, the server can be used locally with Claude Desktop by configuring it in the MCP servers configuration.
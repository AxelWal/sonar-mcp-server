# Sonar MVP Server

An MCP server implementation for the SonarQube Cloud API in Golang.

## Build and Release

```bash
# to quickly build the latest snapshot
goreleaser build --snapshot --clean
goreleaser release --skip=publish --snapshot --clean
```

## Usage Instructions

If you want to use the tool locally, e.g. with Claude Desktop, use the following
configuration for the MCP server.

```json
{
    "mcpServers": {
      "sonar": {
        "command": "/Users/mario-leander.reimer/Applications/sonar-mcp-server",
        "args": ["-t", "stdio"],
        "env": {
          "SONAR_TOKEN": "<<INSERT TOKEN HERE>>"
        }
      }
    }
}
```

Alternatively, you can use the MCP introspector for easy local development:
```bash
# as stdio binary
npx @modelcontextprotocol/inspector go run main.go

# as SSE server using 
go run main.go --transport sse
npx @modelcontextprotocol/inspector npx mcp-remote@next http://localhost:8080/sse
npx @modelcontextprotocol/inspector
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the MIT open source license, read the `LICENSE`
file for details.

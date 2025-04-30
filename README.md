# Sonar MVP Server

An MCP server implementation for the SonarQube Cloud API in Golang.

## Build and Deploy

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

There is also an HTTP (SSE) mode implemented.

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the MIT open source license, read the `LICENSE`
file for details.

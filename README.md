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

## Deployment

Currently using manual Google Cloud Run deployment. Can either be deployed
directly from source or using the Docker image built on Github.

```bash
# create a new secret for the SONAR_TOKEN
gcloud services enable secretmanager.googleapis.com
print $SONAR_TOKEN | gcloud secrets create sonar-token --data-file=-

# next deploy the local build from source to Cloud Run
gcloud services enable run.googleapis.com cloudbuild.googleapis.com artifactregistry.googleapis.com

gcloud secrets add-iam-policy-binding sonar-token \
  --member=serviceAccount:343509396461-compute@developer.gserviceaccount.com \
  --role=roles/secretmanager.secretAccessor

gcloud run deploy sonar-mcp-server --source=. \
  --region=europe-north1 \
  --port=8080 --allow-unauthenticated \
  --set-secrets=SONAR_TOKEN=sonar-token:latest \
  --set-env-vars=BASE_URL=https://sonar-mcp-server-343509396461.europe-north1.run.app
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the MIT open source license, read the `LICENSE`
file for details.

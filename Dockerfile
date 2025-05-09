FROM golang:1.24.3-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sonar-mcp-server -ldflags "-s -w -X main.version=$(date +%Y-%m-%dT%H:%M:%S%z)"

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/sonar-mcp-server /sonar-mcp-server

ENV SONAR_TOKEN=
EXPOSE 8080
CMD ["/sonar-mcp-server", "-t", "sse", "-p", "8080"]

FROM gcr.io/distroless/static-debian12
COPY dist/sonar-mcp-server_linux_amd64_v1/sonar-mcp-server /
CMD ["/sonar-mcp-server"]
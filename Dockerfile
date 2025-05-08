FROM gcr.io/distroless/static-debian12

COPY dist/sonar-mcp-server_linux_amd64_v1/sonar-mcp-server /

ENV SONAR_TOKEN=
EXPOSE 8080
CMD ["/sonar-mcp-server", "-t", "sse", "-p", "8080"]
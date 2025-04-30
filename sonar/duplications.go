package sonar

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func AddDuplicationsTool(s *server.MCPServer) {
	// Create a new MCP tool for listing Sonar projects
	_ = mcp.NewTool("sonar_duplications",
		mcp.WithDescription("Get all duplications between source files within the codebase a Sonar cloud project."),
		mcp.WithString("organization",
			mcp.Required(),
			mcp.Description("The Sonar cloud organization name"),
		),
	)

}

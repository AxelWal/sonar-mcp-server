package sonar

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type DuplicationBlock struct {
	From int    `json:"from"`
	Size int    `json:"size"`
	Ref  string `json:"_ref"`
}

type Duplication struct {
	Blocks []DuplicationBlock `json:"blocks"`
}

type File struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	ProjectName string `json:"projectName"`
}

type DuplicationsResponse struct {
	Duplications []Duplication   `json:"duplications"`
	Files        map[string]File `json:"files"`
}

func AddDuplicationsTool(s *server.MCPServer) {
	// Create a new MCP tool for listing Sonar projects
	duplicationsTool := mcp.NewTool("sonar_duplications",
		mcp.WithDescription("Get duplications between source files, either within a branch or pull request or for a file in the project."),
		mcp.WithString("branch",
			mcp.Description("The SCM branch key or name (optional), e.g. feature/my_branch"),
		),
		mcp.WithString("key",
			// we might need to split the key into project and file
			mcp.Description("The file key (optional), e.g. my_project:/src/foo/Bar.php"),
		),
		mcp.WithString("pull_request",
			mcp.Description("The pull request key (optional), e.g. 5461"),
		),
	)

	// Add the tool to the server
	s.AddTool(duplicationsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract the parameters from the request
		branch := request.Params.Arguments["branch"].(string)
		key := request.Params.Arguments["key"].(string)
		pullRequest := request.Params.Arguments["pull_request"].(string)

		// Call the Sonarcloud API to get the projects
		duplications, err := getDuplications(branch, key, pullRequest)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve duplications.", err), nil
		}

		// Return the projects as a result
		return mcp.NewToolResultText(duplications), nil
	})
}

func getDuplications(branch, key, pullRequest string) (string, error) {
	panic("unimplemented")
}

package sonar

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ProjectsResponse struct {
	Paging     Paging     `json:"paging"`
	Components []Projects `json:"components"`
}

func AddProjectsTool(s *server.MCPServer) {
	// create a new MCP tool for listing Sonar projects
	projectsTool := mcp.NewTool("sonar_projects",
		mcp.WithDescription("List all Sonar projects for a given organization."),
		mcp.WithString("organization",
			mcp.Description("The Sonar cloud organization name, e.g. my_organization."),
			mcp.Required(),
			mcp.DefaultString(""),
		),
	)

	// add the tool to the server
	s.AddTool(projectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in AddProjectsTool: %v\n", r)
			}
		}()
		args := request.GetArguments()
		var organization string
		if v, ok := args["organization"]; ok {
			if s, ok := v.(string); ok {
				organization = s
			}
		}
		projects, err := searchProjects(organization)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve sonar projects.", err), nil
		}
		return mcp.NewToolResultText(projects), nil
	})
}

func searchProjects(organization string) (string, error) {
	url := fmt.Sprintf("https://sonarcloud.io/api/projects/search?organization=%s", organization)

	body, err := performGetRequest(url)
	if err != nil {
		return "", err
	}

	var projectsResponse ProjectsResponse
	err = json.Unmarshal(body, &projectsResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return prettyPrint(projectsResponse.Components)
}

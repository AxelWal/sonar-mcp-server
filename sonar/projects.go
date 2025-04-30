package sonar

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Paging struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

type Projects struct {
	Organization     string `json:"organization"`
	Key              string `json:"key"`
	Name             string `json:"name"`
	Qualifier        string `json:"qualifier"`
	Visibility       string `json:"visibility"`
	LastAnalysisDate string `json:"lastAnalysisDate"`
	Revision         string `json:"revision"`
}

type ProjectsResponse struct {
	Paging     Paging     `json:"paging"`
	Components []Projects `json:"components"`
}

func AddProjectsTool(s *server.MCPServer) {
	// Create a new MCP tool for listing Sonar projects
	projectsTool := mcp.NewTool("sonar_projects",
		mcp.WithDescription("List all Sonar cloud projects for a given organization"),
		mcp.WithString("organization",
			mcp.Required(),
			mcp.Description("The Sonar cloud organization name"),
		),
	)

	// Add the tool to the server
	s.AddTool(projectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract the organization name from the request
		organization := request.Params.Arguments["organization"].(string)

		// Call the Sonarcloud API to get the projects
		projects, err := getProjects(organization)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve sonar projects.", err), nil
		}

		// Return the projects as a result
		return mcp.NewToolResultText(projects), nil
	})
}

func getProjects(organization string) (string, error) {
	url := fmt.Sprintf("https://sonarcloud.io/api/projects/search?organization=%s", organization)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set the authorization header using the global sonarToken
	req.SetBasicAuth(getSonarToken(), "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the response body into the SonarProjects type
	var projectsResponse ProjectsResponse
	err = json.Unmarshal(body, &projectsResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return prettyPrint(projectsResponse.Components)
}

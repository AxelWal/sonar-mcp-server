package sonar

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type IssuesResponse struct {
	Paging     Paging      `json:"paging"`
	Issues     []Issue     `json:"issues"`
	Components []Component `json:"components"`
	Rules      []Rule      `json:"rules"`
	Users      []User      `json:"users"`
}

func AddIssuesTool(s *server.MCPServer) {
	// create a new MCP tool for searching Sonar issues
	issuesTool := mcp.NewTool("sonar_issues",
		mcp.WithDescription("Search and get all issues for a specified Sonar project."),
		mcp.WithString("projectKey",
			mcp.Description("Key of the project or application, e.g. my_project."),
			mcp.Required(),
		),
		mcp.WithString("organization",
			mcp.Description("The Sonar cloud organization key or name, e.g. my_organization."),
		),
		mcp.WithArray("componentKeys",
			mcp.Description("Retrieve issues associated to a specific list of components (and all its descendants). A component can be a project, directory or file, e.g. my_project.src"),
		),
		mcp.WithString("branch",
			mcp.Description("The SCM branch key or name (optional), e.g. feature/my_branch"),
			mcp.DefaultString("main"),
		),
		mcp.WithString("pullRequest",
			mcp.Description("The pull request key (optional), e.g. 5461"),
		),
	)

	// add the tool to the server
	s.AddTool(issuesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// extract the parameters from the request
		projectKey := request.Params.Arguments["projectKey"].(string)
		organization := request.Params.Arguments["organization"].(string)
		componentKeys := request.Params.Arguments["componentKeys"].([]string)
		branch := request.Params.Arguments["branch"].(string)
		pullRequest := request.Params.Arguments["pullRequest"].(string)

		// call the Sonarcloud API to get the issues
		issues, err := searchIssues(organization, projectKey, componentKeys, branch, pullRequest)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve issues.", err), nil
		}

		return mcp.NewToolResultText(issues), nil
	})
}

func searchIssues(organization string, projectKey string, componentKeys []string, branch string, pullRequest string) (string, error) {
	organizationParam := ""
	if organization != "" {
		organizationParam = fmt.Sprintf("&organization=%s", organization)
	}
	componentKeysParam := ""
	if len(componentKeys) > 0 {
		componentKeysParam = fmt.Sprintf("&componentKeys=%s", strings.Join(componentKeys, ","))
	}
	branchParam := ""
	if branch != "" {
		branchParam = fmt.Sprintf("&branch=%s", branch)
	}
	pullRequestParam := ""
	if pullRequest != "" {
		pullRequestParam = fmt.Sprintf("&pullRequest=%s", pullRequest)
	}

	url := fmt.Sprintf("https://sonarcloud.io/api/issues/search?projectKey=%s%s%s%s%s", projectKey, organizationParam, componentKeysParam, branchParam, pullRequestParam)

	body, err := performGetRequest(url)
	if err != nil {
		return "", err
	}

	var response IssuesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return prettyPrint(response)
}

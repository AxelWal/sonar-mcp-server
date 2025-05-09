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
	Issues     []Issue     `json:"issues,omitempty"`
	Components []Component `json:"components,omitempty"`
	Rules      []Rule      `json:"rules,omitempty"`
	Users      []User      `json:"users,omitempty"`
}

func AddIssuesTool(s *server.MCPServer) {
	// create a new MCP tool for searching Sonar issues
	issuesTool := mcp.NewTool("sonar_issues",
		mcp.WithDescription("Search and get all issues for a specified Sonar project."),
		mcp.WithString("projectKey",
			mcp.Description("Key of the project or application, e.g. my_project."),
			mcp.DefaultString(""),
			mcp.Required(),
		),
		mcp.WithString("organization",
			mcp.Description("The Sonar cloud organization key or name, e.g. my_organization."),
			mcp.DefaultString(""),
		),
		mcp.WithString("branch",
			mcp.Description("The SCM branch key or name (optional), e.g. feature/my_branch"),
			mcp.DefaultString("main"),
		),
		mcp.WithArray("impactSeverities",
			mcp.Description("The severity of the issues to be retrieved. Possible values: BLOCKER, HIGH, MEDIUM, LOW, INFO."),
			mcp.DefaultArray([]string{"BLOCKER", "HIGH"}),
			mcp.Enum("BLOCKER", "HIGH", "MEDIUM", "LOW", "INFO"),
		),
		mcp.WithArray("issueStatus",
			mcp.Description("The status of the issues to be retrieved. Possible values: OPEN, CONFIRMED, FALSE_POSITIVE, ACCEPTED, FIXED."),
			mcp.DefaultArray([]string{"OPEN"}),
			mcp.Enum("OPEN", "CONFIRMED", "FALSE_POSITIVE", "ACCEPTED", "FIXED"),
		),
		mcp.WithString("resolved",
			mcp.Description("The resolved status of the issues to be retrieved. Possible values: true, false, yes, no."),
			mcp.DefaultString(""),
			mcp.Enum("true", "false", "yes", "no"),
		),
	)

	// add the tool to the server
	s.AddTool(issuesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// extract the parameters from the request
		projectKey := request.Params.Arguments["projectKey"].(string)
		organization := request.Params.Arguments["organization"].(string)
		branch := request.Params.Arguments["branch"].(string)
		issueStatus := request.Params.Arguments["issueStatus"].([]interface{})
		impactSeverities := request.Params.Arguments["impactSeverities"].([]interface{})
		resolved := request.Params.Arguments["resolved"].(string)

		// call the Sonarcloud API to get the issues
		issues, err := searchIssues(organization, projectKey, branch, issueStatus, resolved, impactSeverities)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve issues.", err), nil
		}

		return mcp.NewToolResultText(issues), nil
	})
}

func searchIssues(organization string, projectKey string, branch string, issueStatus []interface{}, resolved string, impactSeverities []interface{}) (string, error) {
	organizationParam := ""
	if organization != "" {
		organizationParam = fmt.Sprintf("&organization=%s", organization)
	}
	branchParam := ""
	if branch != "" {
		branchParam = fmt.Sprintf("&branch=%s", branch)
	}
	issueStatusParam := ""
	if len(issueStatus) > 0 {
		// join the issue statuses with commas
		issueStatusParam = fmt.Sprintf("&issueStatuses=%s", strings.Join(toStringArray(issueStatus), ","))
	}
	resolvedParam := ""
	if resolved != "" {
		resolvedParam = fmt.Sprintf("&resolved=%s", resolved)
	}
	impactSeveritiesParam := ""
	if len(impactSeverities) > 0 {
		// join the impact severities with commas
		impactSeveritiesParam = fmt.Sprintf("&impactSeverities=%s", strings.Join(toStringArray(impactSeverities), ","))
	}

	// construct the URL for the Sonarcloud API
	url := fmt.Sprintf("https://sonarcloud.io/api/issues/search?projectKey=%s%s%s%s%s%s",
		projectKey, organizationParam, branchParam, issueStatusParam, resolvedParam, impactSeveritiesParam)

	body, err := performGetRequest(url)
	if err != nil {
		return "", err
	}

	var response IssuesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// check if the response contains issues
	if len(response.Issues) == 0 {
		return "No issues found.", nil
	}

	return prettyPrint(response.Issues)
}

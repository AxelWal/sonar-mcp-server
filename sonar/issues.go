package sonar

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("Adding issues tool")
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
		mcp.WithString("pullRequest",
			mcp.Description("The pull request ID or key (optional), e.g. 42 or PR-123"),
			mcp.DefaultString(""),
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
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered in AddIssuesTool: %v", r)
			}
		}()
		// extract the parameters from the request
		args := request.GetArguments()

		var projectKey, organization, branch, pullRequest, resolved string
		var issueStatus, impactSeverities []interface{}

		if v, ok := args["projectKey"]; ok {
			if s, ok := v.(string); ok {
				projectKey = s
			}
		}

		if v, ok := args["organization"]; ok {
			if s, ok := v.(string); ok {
				organization = s
			}
		}

		if v, ok := args["branch"]; ok {
			if s, ok := v.(string); ok {
				branch = s
			}
		}

		if v, ok := args["pullRequest"]; ok {
			if s, ok := v.(string); ok {
				pullRequest = s
			}
		}

		if v, ok := args["issueStatus"]; ok {
			if arr, ok := v.([]interface{}); ok {
				issueStatus = arr
			}
		}

		if v, ok := args["impactSeverities"]; ok {
			if arr, ok := v.([]interface{}); ok {
				impactSeverities = arr
			}
		}

		if v, ok := args["resolved"]; ok {
			if s, ok := v.(string); ok {
				resolved = s
			}
		}

		// call the Sonarcloud API to get the issues
		issues, err := searchIssues(organization, projectKey, branch, pullRequest, issueStatus, resolved, impactSeverities)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve issues.", err), nil
		}

		return mcp.NewToolResultText(issues), nil
	})
}

func searchIssues(organization string, projectKey string, branch string, pullRequest string, issueStatus []interface{}, resolved string, impactSeverities []interface{}) (string, error) {
	organizationParam := ""
	if organization != "" {
		organizationParam = fmt.Sprintf("&organization=%s", organization)
	}
	branchParam := ""
	if branch != "" {
		branchParam = fmt.Sprintf("&branch=%s", branch)
	}
	pullRequestParam := ""
	if pullRequest != "" {
		pullRequestParam = fmt.Sprintf("&pullRequest=%s", pullRequest)
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
	url := fmt.Sprintf("https://sonarcloud.io/api/issues/search?projectKeys=%s%s%s%s%s%s%s",
		projectKey, organizationParam, branchParam, pullRequestParam, issueStatusParam, resolvedParam, impactSeveritiesParam)
	log.Println(url)
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

package sonar

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type DuplicationsResponse struct {
	Duplications []Duplication   `json:"duplications"`
	Files        map[string]File `json:"files"`
}

func AddDuplicationsTool(s *server.MCPServer) {
	// create a new MCP tool for showing duplications
	duplicationsTool := mcp.NewTool("sonar_duplications",
		mcp.WithDescription("Show duplications between source files, either within a branch or pull request or for a file in a Sonar project."),
		mcp.WithString("branch",
			mcp.Description("The SCM branch key or name (optional), e.g. feature/my_branch"),
			mcp.DefaultString("main"),
		),
		mcp.WithString("key",
			// we might need to split the key into project and file
			mcp.Description("The file key (optional), e.g. my_project:/src/foo/Bar.php"),
			mcp.DefaultString(""),
		),
		mcp.WithString("pullRequest",
			mcp.Description("The pull request key (optional), e.g. 5461"),
			mcp.DefaultString(""),
		),
	)

	// add the tool to the server
	s.AddTool(duplicationsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in AddDuplicationsTool: %v\n", r)
			}
		}()
		args := request.GetArguments()
		var branch, key, pullRequest string
		if v, ok := args["branch"]; ok {
			if s, ok := v.(string); ok {
				branch = s
			}
		}
		if v, ok := args["key"]; ok {
			if s, ok := v.(string); ok {
				key = s
			}
		}
		if v, ok := args["pullRequest"]; ok {
			if s, ok := v.(string); ok {
				pullRequest = s
			}
		}
		// call the Sonarcloud API to get the duplications
		duplications, err := showDuplications(branch, key, pullRequest)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve duplications.", err), nil
		}
		return mcp.NewToolResultText(duplications), nil
	})
}

func showDuplications(branch, key, pullRequest string) (string, error) {
	keyParam := ""
	if key != "" {
		keyParam = fmt.Sprintf("&key=%s", key)
	}
	pullRequestParam := ""
	if pullRequest != "" {
		pullRequestParam = fmt.Sprintf("&pullRequest=%s", pullRequest)
	}

	url := fmt.Sprintf("https://sonarcloud.io/api/duplications/show?branch=%s%s%s", branch, keyParam, pullRequestParam)

	body, err := performGetRequest(url)
	if err != nil {
		return "", err
	}

	var response DuplicationsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return prettyPrint(response)
}

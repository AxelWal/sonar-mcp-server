package sonar

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func AddMetricsTool(s *server.MCPServer) {
	metricsTool := mcp.NewTool("sonar_metrics",
		mcp.WithDescription("Get quality metrics for a Sonar project."),
		mcp.WithString("projectKey",
			mcp.Description("Key of the project to get metrics for."),
			mcp.Required(),
		),
		mcp.WithString("branch",
			mcp.Description("The SCM branch (optional)."),
			mcp.DefaultString("main"),
		),
		mcp.WithArray("metricKeys",
			mcp.Description("List of metrics to retrieve (e.g., complexity, ncloc)."),
			mcp.DefaultArray([]string{"ncloc", "complexity", "violations"}),
			mcp.Enum("ncloc", "complexity", "violations"),
		),
	)

	s.AddTool(metricsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in AddMetricsTool: %v\n", r)
			}
		}()
		args := request.GetArguments()
		var projectKey, branch string
		var metricKeys []interface{}
		if v, ok := args["projectKey"]; ok {
			if s, ok := v.(string); ok {
				projectKey = s
			}
		}
		if v, ok := args["branch"]; ok {
			if s, ok := v.(string); ok {
				branch = s
			}
		}
		if v, ok := args["metricKeys"]; ok {
			if arr, ok := v.([]interface{}); ok {
				metricKeys = arr
			}
		}
		metrics, err := getProjectMetrics(projectKey, branch, metricKeys)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve metrics.", err), nil
		}
		return mcp.NewToolResultText(metrics), nil
	})
}

func getProjectMetrics(projectKey, branch string, metricKeys []interface{}) (string, error) {
	branchParam := ""
	if branch != "" {
		branchParam = fmt.Sprintf("&branch=%s", branch)
	}

	metricsParam := strings.Join(toStringArray(metricKeys), ",")

	url := fmt.Sprintf("https://sonarcloud.io/api/measures/component?component=%s&metricKeys=%s%s",
		projectKey, metricsParam, branchParam)

	body, err := performGetRequest(url)
	if err != nil {
		return "", err
	}

	var response MeasuresResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(response.Component.Measures) == 0 {
		return "No metrics found.", nil
	}

	return prettyPrint(response.Component.Measures)
}

package sonar

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type SecurityHotspotsResponse struct {
	Paging     Paging      `json:"paging"`
	Hotspots   []Hotspot   `json:"hotspots"`
	Components []Component `json:"components"`
}

func AddHotspotsTool(s *server.MCPServer) {
	// create a new MCP tool for searching security hotspots
	hotspotsTool := mcp.NewTool("sonar_hotspots",
		mcp.WithDescription("Search and get security hotpots in the source files of a specified Sonar project."),
		mcp.WithString("projectKey",
			mcp.Description("Key of the project or application, e.g. my_project."),
			mcp.Required(),
		),
		mcp.WithArray("files",
			mcp.Description("Array or list of file paths. Returns only hotspots found in those files, e.g. src/foo/Bar.php. This parameter is optional."),
			mcp.DefaultArray([]string{}),
		),
		mcp.WithString("status",
			mcp.Description("The status of the security hotspot, only these are returned, e.g. TO_REVIEW, REVIEWED. This parameter is optional."),
			mcp.DefaultString(""),
			mcp.Enum("TO_REVIEW", "REVIEWED"),
		),
	)

	// add the tool to the server
	s.AddTool(hotspotsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		projectKey := args["projectKey"].(string)
		files := args["files"].([]interface{})
		status := args["status"].(string)

		// call the Sonarcloud API to get the hotspots
		duplications, err := searchHotspots(projectKey, files, status)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to retrieve security hotspots.", err), nil
		}

		return mcp.NewToolResultText(duplications), nil
	})
}

func searchHotspots(projectKey string, files []interface{}, status string) (string, error) {
	filesParam := ""
	if len(files) > 0 {
		filesParam = fmt.Sprintf("&files=%s", strings.Join(toStringArray(files), ","))
	}
	statusParam := ""
	if status != "" {
		statusParam = fmt.Sprintf("&status=%s", status)
	}

	url := fmt.Sprintf("https://sonarcloud.io/api/hotspots/search?projectKey=%s%s%s", projectKey, filesParam, statusParam)

	body, err := performGetRequest(url)
	if err != nil {
		return "", err
	}

	var response SecurityHotspotsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return prettyPrint(response)
}

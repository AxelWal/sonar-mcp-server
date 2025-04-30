package sonar

import "github.com/mark3labs/mcp-go/server"

type IssuesResponse struct {
	Paging     Paging      `json:"paging"`
	Issues     []Issue     `json:"issues"`
	Components []Component `json:"components"`
	Rules      []Rule      `json:"rules"`
	Users      []User      `json:"users"`
}

func AddIssuesTool(s *server.MCPServer) {
	panic("unimplemented")
}

package sonar

import "github.com/mark3labs/mcp-go/server"

type TextRange struct {
	StartLine   int `json:"startLine"`
	EndLine     int `json:"endLine"`
	StartOffset int `json:"startOffset"`
	EndOffset   int `json:"endOffset"`
}

type Hotspot struct {
	Key                      string    `json:"key"`
	Component                string    `json:"component"`
	Project                  string    `json:"project"`
	SecurityCategory         string    `json:"securityCategory"`
	VulnerabilityProbability string    `json:"vulnerabilityProbability"`
	Status                   string    `json:"status"`
	Line                     int       `json:"line"`
	Message                  string    `json:"message"`
	Assignee                 string    `json:"assignee"`
	Author                   string    `json:"author"`
	CreationDate             string    `json:"creationDate"`
	UpdateDate               string    `json:"updateDate"`
	TextRange                TextRange `json:"textRange"`
	RuleKey                  string    `json:"ruleKey"`
}

type Component struct {
	Organization string `json:"organization"`
	Key          string `json:"key"`
	Qualifier    string `json:"qualifier"`
	Name         string `json:"name"`
	LongName     string `json:"longName"`
	Path         string `json:"path"`
}

type SecurityHotspotsResponse struct {
	Paging     Paging      `json:"paging"`
	Hotspots   []Hotspot   `json:"hotspots"`
	Components []Component `json:"components"`
}

func AddHotspotsTool(s *server.MCPServer) {
	panic("unimplemented")
}

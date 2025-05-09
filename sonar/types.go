package sonar

type Projects struct {
	Organization     string `json:"organization"`
	Key              string `json:"key"`
	Name             string `json:"name"`
	Qualifier        string `json:"qualifier"`
	Visibility       string `json:"visibility"`
	LastAnalysisDate string `json:"lastAnalysisDate"`
	Revision         string `json:"revision"`
}

type Paging struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

type DuplicationBlock struct {
	From int    `json:"from"`
	Size int    `json:"size"`
	Ref  string `json:"_ref"`
}

type Duplication struct {
	Blocks []DuplicationBlock `json:"blocks"`
}

type File struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	ProjectName string `json:"projectName"`
}

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
	Enabled      bool   `json:"enabled"`
	Qualifier    string `json:"qualifier"`
	Name         string `json:"name"`
	LongName     string `json:"longName"`
	Path         string `json:"path"`
}

type Issue struct {
	Key                        string            `json:"key"`
	Component                  string            `json:"component"`
	Project                    string            `json:"project"`
	Rule                       string            `json:"rule"`
	IssueStatus                string            `json:"issueStatus"`
	Status                     string            `json:"status"`
	Resolution                 string            `json:"resolution"`
	Severity                   string            `json:"severity"`
	Message                    string            `json:"message"`
	Line                       int               `json:"line"`
	Hash                       string            `json:"hash"`
	Author                     string            `json:"author"`
	Effort                     string            `json:"effort"`
	CreationDate               string            `json:"creationDate"`
	UpdateDate                 string            `json:"updateDate"`
	Tags                       []string          `json:"tags"`
	Type                       string            `json:"type"`
	Comments                   []Comment         `json:"comments"`
	Attr                       map[string]string `json:"attr"`
	Transitions                []string          `json:"transitions"`
	Actions                    []string          `json:"actions"`
	TextRange                  TextRange         `json:"textRange"`
	Flows                      []Flow            `json:"flows"`
	RuleDescriptionContextKey  string            `json:"ruleDescriptionContextKey"`
	CleanCodeAttributeCategory string            `json:"cleanCodeAttributeCategory"`
	CleanCodeAttribute         string            `json:"cleanCodeAttribute"`
	Impacts                    []Impact          `json:"impacts"`
}

type Comment struct {
	Key       string `json:"key"`
	Login     string `json:"login"`
	HtmlText  string `json:"htmlText"`
	Markdown  string `json:"markdown"`
	Updatable bool   `json:"updatable"`
	CreatedAt string `json:"createdAt"`
}

type Flow struct {
	Locations []Location `json:"locations"`
}

type Location struct {
	TextRange TextRange `json:"textRange"`
	Msg       string    `json:"msg"`
}

type Impact struct {
	SoftwareQuality string `json:"softwareQuality"`
	Severity        string `json:"severity"`
}

type Rule struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Lang     string `json:"lang"`
	LangName string `json:"langName"`
}

type User struct {
	Login  string `json:"login"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Avatar string `json:"avatar"`
}

type Measure struct {
	Metric    string  `json:"metric"`
	Value     string  `json:"value"`
	BestValue bool    `json:"bestValue,omitempty"`
	Period    *Period `json:"period,omitempty"`
}

type Period struct {
	Value     string `json:"value"`
	BestValue bool   `json:"bestValue,omitempty"`
}

type MeasuresResponse struct {
	Component struct {
		ID        string    `json:"id"`
		Key       string    `json:"key"`
		Name      string    `json:"name"`
		Qualifier string    `json:"qualifier"`
		Measures  []Measure `json:"measures"`
	} `json:"component"`
}

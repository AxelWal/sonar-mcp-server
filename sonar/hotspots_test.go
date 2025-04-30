package sonar

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalSecurityHotspotsResponse(t *testing.T) {
	jsonData := `{
		"paging": {
			"pageIndex": 1,
			"pageSize": 100,
			"total": 3
		},
		"hotspots": [
			{
				"key": "hotspot-0",
				"component": "com.sonarsource:test-project:src/main/java/com/sonarsource/FourthClass.java",
				"project": "com.sonarsource:test-project",
				"securityCategory": "others",
				"vulnerabilityProbability": "LOW",
				"status": "TO_REVIEW",
				"line": 10,
				"message": "message-0",
				"assignee": "assignee-uuid",
				"author": "joe",
				"creationDate": "2020-01-02T15:43:10+0100",
				"updateDate": "2020-01-02T15:43:10+0100",
				"textRange": {
					"startLine": 2,
					"endLine": 2,
					"startOffset": 0,
					"endOffset": 204
				},
				"ruleKey": "repository:rule-key"
			}
		],
		"components": [
			{
				"organization": "default-organization",
				"key": "com.sonarsource:test-project:src/main/java/com/sonarsource/FourthClass.java",
				"qualifier": "FIL",
				"name": "FourthClass.java",
				"longName": "src/main/java/com/sonarsource/FourthClass.java",
				"path": "src/main/java/com/sonarsource/FourthClass.java"
			}
		]
	}`

	var response SecurityHotspotsResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	assert.NoError(t, err)

	// Verify Paging
	assert.Equal(t, 1, response.Paging.PageIndex)
	assert.Equal(t, 100, response.Paging.PageSize)
	assert.Equal(t, 3, response.Paging.Total)

	// Verify Hotspots
	assert.Len(t, response.Hotspots, 1)
	hotspot := response.Hotspots[0]
	assert.Equal(t, "hotspot-0", hotspot.Key)
	assert.Equal(t, "com.sonarsource:test-project:src/main/java/com/sonarsource/FourthClass.java", hotspot.Component)
	assert.Equal(t, "com.sonarsource:test-project", hotspot.Project)
	assert.Equal(t, "others", hotspot.SecurityCategory)
	assert.Equal(t, "LOW", hotspot.VulnerabilityProbability)
	assert.Equal(t, "TO_REVIEW", hotspot.Status)
	assert.Equal(t, 10, hotspot.Line)
	assert.Equal(t, "message-0", hotspot.Message)
	assert.Equal(t, "assignee-uuid", hotspot.Assignee)
	assert.Equal(t, "joe", hotspot.Author)
	assert.Equal(t, "2020-01-02T15:43:10+0100", hotspot.CreationDate)
	assert.Equal(t, "2020-01-02T15:43:10+0100", hotspot.UpdateDate)
	assert.Equal(t, 2, hotspot.TextRange.StartLine)
	assert.Equal(t, 2, hotspot.TextRange.EndLine)
	assert.Equal(t, 0, hotspot.TextRange.StartOffset)
	assert.Equal(t, 204, hotspot.TextRange.EndOffset)
	assert.Equal(t, "repository:rule-key", hotspot.RuleKey)

	// Verify Components
	assert.Len(t, response.Components, 1)
	component := response.Components[0]
	assert.Equal(t, "default-organization", component.Organization)
	assert.Equal(t, "com.sonarsource:test-project:src/main/java/com/sonarsource/FourthClass.java", component.Key)
	assert.Equal(t, "FIL", component.Qualifier)
	assert.Equal(t, "FourthClass.java", component.Name)
	assert.Equal(t, "src/main/java/com/sonarsource/FourthClass.java", component.LongName)
	assert.Equal(t, "src/main/java/com/sonarsource/FourthClass.java", component.Path)
}
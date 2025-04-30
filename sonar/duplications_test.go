package sonar

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicationsResponseUnmarshal(t *testing.T) {
	jsonData := `{
		"duplications": [
			{
				"blocks": [
					{"from": 94, "size": 101, "_ref": "1"},
					{"from": 83, "size": 101, "_ref": "2"}
				]
			},
			{
				"blocks": [
					{"from": 38, "size": 40, "_ref": "1"},
					{"from": 29, "size": 39, "_ref": "2"}
				]
			}
		],
		"files": {
			"1": {
				"key": "org.codehaus.sonar:sonar-plugin-api:src/main/java/org/sonar/api/utils/command/CommandExecutor.java",
				"name": "CommandExecutor",
				"projectName": "SonarQube"
			},
			"2": {
				"key": "com.sonarsource.orchestrator:sonar-orchestrator:src/main/java/com/sonar/orchestrator/util/CommandExecutor.java",
				"name": "CommandExecutor",
				"projectName": "SonarSource :: Orchestrator"
			}
		}
	}`

	var response DuplicationsResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	assert.NoError(t, err)

	assert.Len(t, response.Duplications, 2)
	assert.Len(t, response.Duplications[0].Blocks, 2)
	assert.Equal(t, 94, response.Duplications[0].Blocks[0].From)
	assert.Equal(t, "1", response.Duplications[0].Blocks[0].Ref)

	assert.Len(t, response.Files, 2)
	assert.Equal(t, "SonarQube", response.Files["1"].ProjectName)
	assert.Equal(t, "CommandExecutor", response.Files["2"].Name)
}

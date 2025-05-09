package sonar

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMeasuresResponse(t *testing.T) {
	jsonData := `{
		"component": {
			"id": "AU_3HQlfS96Al-sWkzB0",
			"key": "my-project",
			"name": "My Project",
			"qualifier": "TRK",
			"measures": [
				{
					"metric": "complexity",
					"value": "4214"
				},
				{
					"metric": "code_smells",
					"value": "8595",
					"bestValue": false
				},
				{
					"metric": "ncloc",
					"value": "51667"
				},
				{
					"metric": "bugs",
					"value": "12",
					"period": {
						"value": "2",
						"bestValue": false
					}
				}
			]
		}
	}`

	var response MeasuresResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	assert.NoError(t, err, "Failed to unmarshal JSON")
	assert.Equal(t, "my-project", response.Component.Key, "Expected project key to be 'my-project'")
	assert.Len(t, response.Component.Measures, 4, "Expected 4 measures")
	
	// Check specific metrics
	assert.Equal(t, "complexity", response.Component.Measures[0].Metric, "Expected first metric to be 'complexity'")
	assert.Equal(t, "4214", response.Component.Measures[0].Value, "Expected complexity value to be '4214'")
	
	assert.Equal(t, "code_smells", response.Component.Measures[1].Metric, "Expected second metric to be 'code_smells'")
	assert.Equal(t, false, response.Component.Measures[1].BestValue, "Expected code_smells bestValue to be false")
	
	// Check period data
	assert.Equal(t, "bugs", response.Component.Measures[3].Metric, "Expected fourth metric to be 'bugs'")
	assert.NotNil(t, response.Component.Measures[3].Period, "Expected period data to be present")
	assert.Equal(t, "2", response.Component.Measures[3].Period.Value, "Expected period value to be '2'")
}
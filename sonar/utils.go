package sonar

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(data any) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %w", err)
	}

	// Return the pretty-printed JSON as a string
	return string(jsonData), nil
}

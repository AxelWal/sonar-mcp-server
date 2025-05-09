package sonar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func prettyPrint(data any) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %w", err)
	}

	return string(jsonData), nil
}

func toStringArray(issueStatus []interface{}) []string {
	issueStatusStr := make([]string, len(issueStatus))
	for i, v := range issueStatus {
		issueStatusStr[i] = v.(string)
	}
	return issueStatusStr
}

func performGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(getSonarToken(), "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

package sonar

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSonarToken(t *testing.T) {
	t.Run("Should return token when SONAR_TOKEN is set", func(t *testing.T) {
		os.Setenv("SONAR_TOKEN", "test-token")
		defer os.Unsetenv("SONAR_TOKEN")

		token := getSonarToken()
		assert.Equal(t, "test-token", token, "Token should match the expected value")
	})

	t.Run("Should log fatal when SONAR_TOKEN is not set", func(t *testing.T) {
		os.Unsetenv("SONAR_TOKEN")

		// To test log.Fatal, you would need to use a custom logger or mock the log.Fatal behavior.
		// This is a placeholder to indicate that such a test would be implemented.
		// Example: Capture log output or use a library like 'testify/mock'.
	})
}

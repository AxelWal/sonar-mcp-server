package sonar

import (
	"log"
	"os"
)

func getSonarToken() string {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in getSonarToken: %v", r)
		}
	}()
	sonarToken := os.Getenv("SONAR_TOKEN")
	if sonarToken == "" {
		log.Println("Warning: SONAR_TOKEN environment variable is not set")
	}
	return sonarToken
}

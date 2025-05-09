package sonar

import (
	"log"
	"os"
)

func getSonarToken() string {
	sonarToken := os.Getenv("SONAR_TOKEN")
	if sonarToken == "" {
		log.Fatal("SONAR_TOKEN environment variable is not set")
	}
	log.Printf("Using Sonar token: %s", sonarToken)
	return sonarToken
}

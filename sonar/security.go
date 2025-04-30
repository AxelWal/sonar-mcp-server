package sonar

import (
	"log"
	"os"
)

var SonarToken string

func init() {
	SonarToken = os.Getenv("SONAR_TOKEN")
	if SonarToken == "" {
		log.Fatal("SONAR_TOKEN environment variable is not set")
	}
}

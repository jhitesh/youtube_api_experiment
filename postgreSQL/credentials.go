package postgreSQL

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DBCredentials struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var dbCredentials = &DBCredentials{}

const dbCredentialsFile = "database_credentials.yaml"

func readCredentials() {
	yFile, err := os.ReadFile(dbCredentialsFile)
	if err != nil {
		log.Fatal("Error in reading contents of ", dbCredentialsFile, ". ERROR: ", err.Error())
		return
	}

	err = yaml.Unmarshal(yFile, dbCredentials)
	if err != nil {
		log.Fatal("Error in unmarshalling yaml data into queryParams: ", err.Error())
		return
	}
}

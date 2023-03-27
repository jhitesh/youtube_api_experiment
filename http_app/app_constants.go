package http_app

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type AppConstants struct {
	EntriesPerPage int `yaml:"entriesPerPage"`
}

var appConstants = &AppConstants{}
var EntriesPerPage = 0

const constantsFileName = "app_constants.yaml"

func init() {
	yFile, err := os.ReadFile(constantsFileName)
	if err != nil {
		log.Fatal("Error in reading contents of ", constantsFileName, ". ERROR: ", err.Error())
		return
	}

	err = yaml.Unmarshal(yFile, appConstants)
	if err != nil {
		log.Fatal("Error in unmarshalling yaml data into queryParams: ", err.Error())
		return
	}

	EntriesPerPage = appConstants.EntriesPerPage
}

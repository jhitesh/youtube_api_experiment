package query_params_manager

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const queryParamsFileName = "query_params_values.yaml"

func init() {
	yFile, err := os.ReadFile(queryParamsFileName)
	if err != nil {
		log.Fatal("Error in reading contents of ", queryParamsFileName, ". ERROR: ", err.Error())
		return
	}

	err = yaml.Unmarshal(yFile, queryParams)
	if err != nil {
		log.Fatal("Error in unmarshalling yaml data into queryParams: ", err.Error())
		return
	}

	queryParams.defaultFill()

	QueryManager = &QueryParamsManager{
		QueryParams: queryParams,
	}
}

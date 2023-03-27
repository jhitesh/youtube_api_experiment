package youtube_api

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type WorkerValues struct {
	Hours   int `yaml:"hours"`
	Minutes int `yaml:"minutes"`
	Seconds int `yaml:"seconds"`
}

func (values *WorkerValues) totalDurationInSeconds() int {
	return values.Seconds + values.Minutes*60 + values.Hours*3600
}

const WORKER_VALUES_FILENAME = "youtube_fetch_worker_values.yaml"

var TimeInterval time.Duration

func init() {
	yFile, err := os.ReadFile(WORKER_VALUES_FILENAME)
	if err != nil {
		log.Fatal("Error in reading contents of ", WORKER_VALUES_FILENAME, ". ERROR: ", err.Error())
		return
	}

	workerValues := &WorkerValues{}
	err = yaml.Unmarshal(yFile, workerValues)
	if err != nil {
		log.Fatal("Error in unmarshalling yaml data into queryParams: ", err.Error())
		return
	}

	TimeInterval = getTimeInterval(workerValues)
}

func getTimeInterval(values *WorkerValues) time.Duration {
	currTime := time.Now()

	addedTime := time.Date(currTime.Year(), currTime.Month(), currTime.Day(), currTime.Hour()+values.Hours, currTime.Minute()+values.Minutes, currTime.Second()+values.Seconds, currTime.Nanosecond(), currTime.Location())

	return addedTime.Sub(currTime)
}

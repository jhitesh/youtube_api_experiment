package main

import (
	"log"
	"net/http"
	"youtube_api_experiment/http_app"
	"youtube_api_experiment/iterative_worker"
	"youtube_api_experiment/youtube_api"
)

func main() {
	videoFetchWorker := iterative_worker.Worker{
		Done:         make(chan bool),
		TimeInterval: youtube_api.TimeInterval,
		JobFunction:  youtube_api.SearchVideos,
	}

	go videoFetchWorker.StartWorking()

	http.HandleFunc("/videos/all", http_app.GetAllVideos)
	http.HandleFunc("/videos/search", http_app.SearchVideos)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

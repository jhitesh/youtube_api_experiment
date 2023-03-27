package http_app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"youtube_api_experiment/postgreSQL"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	pageNumber, _ := strconv.Atoi(r.URL.Query().Get("page"))

	videos, err := postgreSQL.DBConn.GetAllVideos(pageNumber, EntriesPerPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(videos)
	if err != nil {
		log.Println("ERROR in encoding data to ResponseWriter. ERROR: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func SearchVideos(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("q")
	pageNumber, _ := strconv.Atoi(r.URL.Query().Get("page"))

	videos := postgreSQL.DBConn.SearchVideos(searchQuery, pageNumber, EntriesPerPage)

	err := json.NewEncoder(w).Encode(videos)
	if err != nil {
		log.Println("ERROR in encoding data to ResponseWriter. ERROR: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

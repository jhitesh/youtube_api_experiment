package youtube_api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"youtube_api_experiment/query_params_manager"
)

const YOUTUBE_SEARCH_API = "https://www.googleapis.com/youtube/v3/search"

func SearchVideos() {
	resp, err := http.Get(fmt.Sprintf("%s?%s", YOUTUBE_SEARCH_API, query_params_manager.QueryManager.GetURLQueryString()))
	if err != nil {
		log.Fatalln("ERROR in getting response from youtube: ", err.Error())
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	searchResult := &searchResultFormat{}
	err = json.NewDecoder(resp.Body).Decode(searchResult)
	if err != nil {
		log.Fatalln("ERROR in decoding response body: ", err.Error())
	}

	searchResult.processAndStore()
}

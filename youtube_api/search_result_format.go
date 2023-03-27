package youtube_api

import (
	"youtube_api_experiment/postgreSQL"
	"youtube_api_experiment/query_params_manager"
)

type searchResultFormat struct {
	Kind          string            `json:"kind"`
	Etag          string            `json:"etag"`
	NextPageToken string            `json:"nextPageToken"`
	PrevPageToken string            `json:"prevPageToken"`
	RegionCode    string            `json:"regionCode"`
	PageInfo      singlePageInfo    `json:"pageInfo"`
	Items         []singleVideoData `json:"items"`
}

func (result *searchResultFormat) processAndStore() {
	query_params_manager.QueryManager.UpdatePageToken(result.NextPageToken)

	for _, videoData := range result.Items {
		videoData.processAndStore()
	}
}

type singlePageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type singleVideoData struct {
	Kind    string             `json:"kind"`
	Etag    string             `json:"etag"`
	Id      videoEntityDetails `json:"id"`
	Snippet videoSnippet       `json:"snippet"`
}

func (data *singleVideoData) processAndStore() {
	dbVideoData := data.toDBVideoData()

	_ = postgreSQL.DBConn.SaveVideoData(dbVideoData)
}

func (data *singleVideoData) toDBVideoData() *postgreSQL.DBVideoData {
	return &postgreSQL.DBVideoData{
		ID:          data.Id.VideoId,
		PublishedAt: data.Snippet.PublishedAt,
		Title:       data.Snippet.Title,
		Description: data.Snippet.Description,
	}
}

type videoEntityDetails struct {
	Kind       string `json:"kind"`
	VideoId    string `json:"videoId"`
	ChannelId  string `json:"channelId"`
	PlaylistId string `json:"playlistId"`
}

type videoSnippet struct {
	PublishedAt          string                            `json:"publishedAt"`
	ChannelId            string                            `json:"channelId"`
	Title                string                            `json:"title"`
	Description          string                            `json:"description"`
	Thumbnails           map[string]singleThumbnailDetails `json:"thumbnails"`
	ChannelTitle         string                            `json:"channelTitle"`
	LiveBroadcastContent string                            `json:"liveBroadcastContent"`
}

type singleThumbnailDetails struct {
	Url    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

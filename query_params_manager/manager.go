package query_params_manager

import (
	"fmt"
	"strings"
)

type QueryParamsManager struct {
	QueryParams *QueryParamsFormat
	UrlQuery    string
}

var QueryManager *QueryParamsManager

func (manager *QueryParamsManager) GetURLQueryString() string {
	if len(manager.UrlQuery) == 0 {
		manager.setUrlQuery()
	}

	return fmt.Sprintf("%s&%s=%s", manager.UrlQuery, "pageToken", manager.QueryParams.PageToken)
}

func (manager *QueryParamsManager) setUrlQuery() {
	individualParams := []string{
		fmt.Sprintf("%s=%s", "key", QueryManager.QueryParams.Key),
		fmt.Sprintf("%s=%s", "part", QueryManager.QueryParams.Part),
		fmt.Sprintf("%s=%s", "maxResults", QueryManager.QueryParams.MaxResults),
		fmt.Sprintf("%s=%s", "order", QueryManager.QueryParams.Order),
		fmt.Sprintf("%s=%s", "publushedAfter", QueryManager.QueryParams.PublishedAfter),
		fmt.Sprintf("%s=%s", "q", QueryManager.QueryParams.Query),
		fmt.Sprintf("%s=%s", "type", QueryManager.QueryParams.Type),
	}

	manager.UrlQuery = strings.Join(individualParams, "&")
}

func (manager *QueryParamsManager) UpdatePageToken(newPageToken string) {
	manager.QueryParams.PageToken = newPageToken
}

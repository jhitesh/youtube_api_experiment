package query_params_manager

import (
	"time"
)

type QueryParamsFormat struct {
	Key            string `yaml:"ApiKey"`
	Part           string `yaml:"part"`
	MaxResults     string `yaml:"maxResults"`
	Order          string `yaml:"order"`
	PublishedAfter string `yaml:"publishedAfter"`
	Query          string `yaml:"Query"`
	Type           string `yaml:"type"`
	PageToken      string `yaml:"pageToken"`
}

type QueryParamsYaml struct {
	Keys  string `yaml:"ApiKeys"`
	Query string `yaml:"Query"`
}

var queryParams = &QueryParamsFormat{}

func (params *QueryParamsFormat) defaultFill() {
	params.Part = "snippet"
	params.MaxResults = "50"
	params.Order = "date"
	params.PublishedAfter = time.Now().AddDate(-1, 0, 0).Format(time.RFC3339)
	params.Type = "video"
}

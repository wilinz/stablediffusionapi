package sdimgedit

import (
	"github.com/go-resty/resty/v2"
	"stablediffusionapi/sdconfig"
)

var (
	httpClient *resty.Client
	apiKey     string
)

func Init(config sdconfig.SDConfig) {
	httpClient = config.HttpClient
	apiKey = config.ApiKey
}

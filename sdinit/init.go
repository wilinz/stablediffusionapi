package sdinit

import (
	"github.com/go-resty/resty/v2"
	"stablediffusionapi/sd3d"
	"stablediffusionapi/sdcommunity"
	"stablediffusionapi/sdconfig"
	"stablediffusionapi/sdimgedit"
	"stablediffusionapi/sdstandard"
)

var currentHttpClient *resty.Client

const baseUrl = "https://stablediffusionapi.com/api/"

func init() {
	config := sdconfig.SDConfig{
		HttpClient: resty.New(),
		BaseUrl:    baseUrl,
	}
	Init(config)
}

func Init(config sdconfig.SDConfig) {
	if config.HttpClient == nil {
		config.HttpClient = currentHttpClient
	} else {
		currentHttpClient = config.HttpClient
	}
	if config.BaseUrl == "" {
		config.BaseUrl = baseUrl
	}
	config.HttpClient.SetBaseURL(config.BaseUrl)
	sd3d.Init(config)
	sdcommunity.Init(config)
	sdimgedit.Init(config)
	sdstandard.Init(config)
}

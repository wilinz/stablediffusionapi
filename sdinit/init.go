package sdinit

import (
	"github.com/go-resty/resty/v2"
	"github.com/wilinz/stablediffusionapi/sd3d"
	"github.com/wilinz/stablediffusionapi/sdcommunity"
	"github.com/wilinz/stablediffusionapi/sdconfig"
	"github.com/wilinz/stablediffusionapi/sdimgedit"
	"github.com/wilinz/stablediffusionapi/sdstandard"
)

var currentHttpClient *resty.Client

const baseUrl = "https://github.com/wilinz/stablediffusionapi.com/api/"

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

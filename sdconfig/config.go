package sdconfig

import "github.com/go-resty/resty/v2"

type SDConfig struct {
	HttpClient *resty.Client
	BaseUrl    string
	ApiKey     string
}

package sdcommunity

import (
	"github.com/go-resty/resty/v2"
	"stablediffusionapi/sdcommunity/model"
)

func ModelReload(payload model.ModelReloadRequest) (*resty.Response, error) {
	url := "v4/dreambooth/model_reload"

	payload.Key = apiKey
	return httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.Text2ImgResponse{}).
		Post(url)
}

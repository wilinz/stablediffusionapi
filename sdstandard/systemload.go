package sdstandard

import (
	"github.com/wilinz/stablediffusionapi/sdstandard/model"
)

func GetSystemLoad() (*model.SystemLoadResponse, error) {
	url := "v3/system_load"

	payload := model.KeyRequest{}
	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.SystemLoadResponse{}).
		Post(url)

	result := resp.Result().(*model.SystemLoadResponse)

	return result, err
}

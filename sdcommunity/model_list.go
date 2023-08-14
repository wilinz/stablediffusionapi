package sdcommunity

import (
	"github.com/wilinz/stablediffusionapi/sdcommunity/model"
)

func ModelList() (*[]model.Model, error) {
	url := "v4/dreambooth/model_list"

	payload := model.KeyRequest{}
	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult([]model.Model{}).
		Post(url)

	result := resp.Result().(*[]model.Model)

	return result, err
}

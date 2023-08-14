package sdimgedit

import "github.com/wilinz/stablediffusionapi/sdimgedit/model"

func Interior(payload model.InteriorRequest) (*model.InteriorResponse, error) {
	url := "v5/interior"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.InteriorResponse{}).
		Post(url)

	result := resp.Result().(*model.InteriorResponse)

	return result, err
}

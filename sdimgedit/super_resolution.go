package sdimgedit

import "stablediffusionapi/sdimgedit/model"

func SuperResolution(payload model.SuperResolutionRequest) (*model.SuperResolutionResponse, error) {
	url := "v3/super_resolution"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.SuperResolutionResponse{}).
		Post(url)

	result := resp.Result().(*model.SuperResolutionResponse)

	return result, err
}

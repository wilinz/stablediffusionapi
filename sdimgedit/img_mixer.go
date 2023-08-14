package sdimgedit

import "stablediffusionapi/sdimgedit/model"

func ImgMixer(payload model.ImgMixerRequest) (*model.ImgMixerResponse, error) {
	url := "v5/interior"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.ImgMixerResponse{}).
		Post(url)

	result := resp.Result().(*model.ImgMixerResponse)

	return result, err
}

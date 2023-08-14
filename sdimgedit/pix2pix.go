package sdimgedit

import "stablediffusionapi/sdimgedit/model"

func Pix2Pix(payload model.Pix2PixRequest) (*model.Pix2PixResponse, error) {
	url := "v5/pix2pix"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.Pix2PixResponse{}).
		Post(url)

	result := resp.Result().(*model.Pix2PixResponse)

	return result, err
}

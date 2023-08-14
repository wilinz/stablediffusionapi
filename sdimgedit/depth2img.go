package sdimgedit

import "stablediffusionapi/sdimgedit/model"

func Depth2Img(payload model.Depth2ImgRequest) (*model.Depth2ImgResponse, error) {
	url := "v5/depth2img"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.Depth2ImgResponse{}).
		Post(url)

	result := resp.Result().(*model.Depth2ImgResponse)

	return result, err
}

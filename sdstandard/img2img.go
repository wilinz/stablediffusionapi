package sdstandard

import (
	model2 "github.com/wilinz/stablediffusionapi/sdstandard/model"
)

func Img2Img(payload model2.Img2ImgRequest) (*model2.Img2ImgResponse, error) {
	url := "v3/img2img"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model2.Img2ImgResponse{}).
		Post(url)

	result := resp.Result().(*model2.Img2ImgResponse)

	return result, err
}

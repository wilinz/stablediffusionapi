package sdstandard

import (
	model2 "github.com/wilinz/stablediffusionapi/sdstandard/model"
)

func Text2Img(payload model2.Text2ImgRequest) (*model2.Text2ImgResponse, error) {
	url := "v3/text2img"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model2.Text2ImgResponse{}).
		Post(url)

	result := resp.Result().(*model2.Text2ImgResponse)

	return result, err
}

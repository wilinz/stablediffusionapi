package sdcommunity

import (
	"stablediffusionapi/sdcommunity/model"
)

func Img2Img(payload model.Img2ImgRequest) (*model.Img2ImgResponse, error) {
	url := "v4/dreambooth/img2img"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.Img2ImgResponse{}).
		Post(url)

	result := resp.Result().(*model.Img2ImgResponse)

	return result, err
}

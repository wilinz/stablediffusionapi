package sdcommunity

import (
	"github.com/wilinz/stablediffusionapi/sdcommunity/model"
)

func Text2Img(payload model.Text2ImgRequest) (*model.Text2ImgResponse, error) {
	url := "v4/dreambooth"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.Text2ImgResponse{}).
		Post(url)

	result := resp.Result().(*model.Text2ImgResponse)

	return result, err
}

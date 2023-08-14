package sdstandard

import (
	"github.com/wilinz/stablediffusionapi/sdstandard/model"
)

func TextToVideo(payload model.TextToVideoRequest) (*model.TextToVideoResponse, error) {
	url := "v5/text2video"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.TextToVideoResponse{}).
		Post(url)

	result := resp.Result().(*model.TextToVideoResponse)

	return result, err
}

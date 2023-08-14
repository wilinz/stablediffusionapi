package sdstandard

import (
	"fmt"
	model2 "stablediffusionapi/sdstandard/model"
)

func FetchQueuedImages(id int64) (*model2.FetchQueuedImagesResponse, error) {
	url := fmt.Sprintf("v3/fetch/%d", id)

	payload := model2.KeyRequest{}
	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model2.FetchQueuedImagesResponse{}).
		Post(url)

	result := resp.Result().(*model2.FetchQueuedImagesResponse)

	return result, err
}

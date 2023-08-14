package sdcommunity

import (
	"github.com/wilinz/stablediffusionapi/sdcommunity/model"
)

func FetchQueuedImages(id int64) (*model.FetchQueuedImagesResponse, error) {
	url := "v4/dreambooth/fetch"

	payload := model.FetchQueuedImagesRequest{}
	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.FetchQueuedImagesResponse{}).
		Post(url)

	result := resp.Result().(*model.FetchQueuedImagesResponse)

	return result, err
}

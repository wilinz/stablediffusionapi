package sdstandard

import (
	"stablediffusionapi/sdstandard/model"
)

func InPaint(payload model.InPaintRequest) (*model.InPaintResponse, error) {
	url := "v3/inpaint"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.InPaintResponse{}).
		Post(url)

	result := resp.Result().(*model.InPaintResponse)

	return result, err
}

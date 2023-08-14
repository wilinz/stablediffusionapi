package sd3d

import "github.com/wilinz/stablediffusionapi/sd3d/model"

func TextTo3d(payload model.TextTo3dRequest) (*model.TextTo3dResponse, error) {
	url := "v3/txt_to_3d"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.TextTo3dResponse{}).
		Post(url)

	result := resp.Result().(*model.TextTo3dResponse)

	return result, err
}

package sd3d

import "github.com/wilinz/stablediffusionapi/sd3d/model"

func ImgTo3d(payload model.ImgTo3dRequest) (*model.ImgTo3dResponse, error) {
	url := "v3/img_to_3d"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.ImgTo3dResponse{}).
		Post(url)

	result := resp.Result().(*model.ImgTo3dResponse)

	return result, err
}

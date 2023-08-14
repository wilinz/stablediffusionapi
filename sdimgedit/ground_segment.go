package sdimgedit

import "stablediffusionapi/sdimgedit/model"

func GroundSegment(payload model.GroundSegmentRequest) (*model.GroundSegmentResponse, error) {
	url := "v5/ground_segment"

	payload.Key = apiKey
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(model.GroundSegmentResponse{}).
		Post(url)

	result := resp.Result().(*model.GroundSegmentResponse)

	return result, err
}

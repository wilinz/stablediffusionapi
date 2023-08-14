package model

type GroundSegmentRequest struct {
	Key         string `json:"key"`
	Prompt      string `json:"prompt"`
	MaskPrompt  string `json:"mask_prompt"`
	ImagePrompt string `json:"image_prompt"`
	Samples     int    `json:"samples"`
}

type GroundSegmentResponse struct {
	Status         string            `json:"status"`
	GenerationTime float64           `json:"generationTime"`
	Id             int               `json:"id"`
	Output         []string          `json:"output"`
	Meta           GroundSegmentMeta `json:"meta"`
}

type GroundSegmentMeta struct {
	ImagePrompt string      `json:"image_prompt"`
	MaskPrompt  string      `json:"mask_prompt"`
	Prompt      string      `json:"prompt"`
	Samples     int         `json:"samples"`
	TrackId     interface{} `json:"track_id"`
	Webhook     interface{} `json:"webhook"`
}

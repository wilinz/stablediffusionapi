package model

type SuperResolutionRequest struct {
	Key         string      `json:"key"`
	Url         string      `json:"url"`
	Scale       int         `json:"scale"`
	Webhook     interface{} `json:"webhook"`
	FaceEnhance bool        `json:"face_enhance"`
}

type SuperResolutionResponse struct {
	Status         string  `json:"status"`
	GenerationTime float64 `json:"generationTime"`
	Id             int     `json:"id"`
	Output         string  `json:"output"`
}

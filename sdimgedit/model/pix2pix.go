package model

type Pix2PixRequest struct {
	Key                string `json:"key"`
	InitImage          string `json:"init_image"`
	Prompt             string `json:"prompt"`
	ImageGuidanceScale int    `json:"image_guidance_scale"`
	Steps              int    `json:"steps"`
	GuidanceScale      int    `json:"guidance_scale"`
}

type Pix2PixResponse struct {
	Status         string  `json:"status"`
	GenerationTime float64 `json:"generationTime"`
	Id             int     `json:"id"`
	Output         string  `json:"output"`
}

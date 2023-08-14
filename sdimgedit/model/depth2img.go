package model

type Depth2ImgRequest struct {
	Key       string `json:"key"`
	InitImage string `json:"init_image"`
	Prompt    string `json:"prompt"`
}

type Depth2ImgResponse struct {
	Status         string  `json:"status"`
	GenerationTime float64 `json:"generationTime"`
	Id             int     `json:"id"`
	Output         string  `json:"output"`
}

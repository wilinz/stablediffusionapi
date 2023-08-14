package model

type TextToVideoRequest struct {
	Key            string `json:"key"`
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
	Scheduler      string `json:"scheduler"`
	Seconds        int    `json:"seconds"`
}

type TextToVideoResponse struct {
	Status         string   `json:"status"`
	GenerationTime float64  `json:"generationTime"`
	Id             int      `json:"id"`
	Output         []string `json:"output"`
}

package model

type ImgTo3dRequest struct {
	Key           string `json:"key"`
	Image         string `json:"image"`
	GuidanceScale int    `json:"guidance_scale"`
	Steps         int    `json:"steps"`
	FrameSize     int    `json:"frame_size"`
	BatchSize     int    `json:"batch_size"`
	Seed          int    `json:"seed"`
	SafetyChecker string `json:"safety_checker"`
	Seconds       int    `json:"seconds"`
	Webhook       string `json:"webhook"`
	TrackId       string `json:"track_id"`
}

type ImgTo3dMeta struct {
	FilePrefix     string `json:"file_prefix"`
	FrameSize      int    `json:"frame_size"`
	GuidanceScale  int    `json:"guidance_scale"`
	NegativePrompt string `json:"negative_prompt"`
	Outdir         string `json:"outdir"`
	OutputType     string `json:"output_type"`
	Prompt         string `json:"prompt"`
	Safetychecker  string `json:"safetychecker"`
	Seconds        int    `json:"seconds"`
	Seed           int    `json:"seed"`
	Steps          int    `json:"steps"`
}
type ImgTo3dResponse struct {
	Status         string      `json:"status"`
	GenerationTime int         `json:"generationTime"`
	ID             int         `json:"id"`
	Output         []string    `json:"output"`
	Meta           ImgTo3dMeta `json:"meta"`
}

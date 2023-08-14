package model

type TextTo3dRequest struct {
	Key           string      `json:"key"`
	Prompt        string      `json:"prompt"`
	GuidanceScale int         `json:"guidance_scale"`
	Steps         int         `json:"steps"`
	FrameSize     int         `json:"frame_size"`
	OutputType    string      `json:"output_type"`
	Webhook       interface{} `json:"webhook"`
	TrackId       interface{} `json:"track_id"`
}

type TextTo3dResponse struct {
	Status         string           `json:"status"`
	GenerationTime int              `json:"generationTime"`
	ID             int              `json:"id"`
	Output         []string         `json:"output"`
	Meta           TextTo3dMetaData `json:"meta"`
}

type TextTo3dMetaData struct {
	FilePrefix     string `json:"file_prefix"`
	FrameSize      int    `json:"frame_size"`
	GuidanceScale  int    `json:"guidance_scale"`
	NegativePrompt string `json:"negative_prompt"`
	OutDir         string `json:"outdir"`
	OutputType     string `json:"output_type"`
	Prompt         string `json:"prompt"`
	SafetyChecker  string `json:"safetychecker"`
	Seconds        int    `json:"seconds"`
	Seed           int    `json:"seed"`
	Steps          int    `json:"steps"`
}

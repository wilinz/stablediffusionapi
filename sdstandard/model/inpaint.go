package model

type InPaintRequest struct {
	Key               string  `json:"key"`
	Prompt            string  `json:"prompt"`
	NegativePrompt    string  `json:"negative_prompt"`
	InitImage         string  `json:"init_image"`
	MaskImage         string  `json:"mask_image"`
	Width             string  `json:"width"`
	Height            string  `json:"height"`
	Samples           string  `json:"samples"`
	NumInferenceSteps string  `json:"num_inference_steps"`
	SafetyChecker     string  `json:"safety_checker"`
	EnhancePrompt     string  `json:"enhance_prompt"`
	GuidanceScale     float64 `json:"guidance_scale"`
	Strength          float64 `json:"strength"`
	Seed              *int64  `json:"seed"`
	Webhook           string  `json:"webhook"`
	TrackId           string  `json:"track_id"`
}

type InPaintResponse struct {
	Status         string         `json:"status"`
	GenerationTime float64        `json:"generationTime"`
	Id             int64          `json:"id"`
	Output         []string       `json:"output"`
	Meta           InPaintingMeta `json:"meta"`
}

type InPaintingMeta struct {
	H              int     `json:"H"`
	W              int     `json:"W"`
	FilePrefix     string  `json:"file_prefix"`
	GuidanceScale  float64 `json:"guidance_scale"`
	InitImage      string  `json:"init_image"`
	MaskImage      string  `json:"mask_image"`
	NSamples       int     `json:"n_samples"`
	NegativePrompt string  `json:"negative_prompt"`
	Outdir         string  `json:"outdir"`
	Prompt         string  `json:"prompt"`
	Safetychecker  string  `json:"safetychecker"`
	Seed           int64   `json:"seed"`
	Steps          int     `json:"steps"`
	Strength       float64 `json:"strength"`
}

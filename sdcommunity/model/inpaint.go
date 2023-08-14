package model

type InPaintRequest struct {
	Key             string      `json:"key"`
	ModelId         string      `json:"model_id"`
	Prompt          string      `json:"prompt"`
	NegativePrompt  string      `json:"negative_prompt"`
	InitImage       string      `json:"init_image"`
	MaskImage       string      `json:"mask_image"`
	Width           string      `json:"width"`
	Height          string      `json:"height"`
	Samples         string      `json:"samples"`
	Steps           string      `json:"steps"`
	SafetyChecker   string      `json:"safety_checker"`
	EnhancePrompt   string      `json:"enhance_prompt"`
	GuidanceScale   float64     `json:"guidance_scale"`
	Strength        float64     `json:"strength"`
	Scheduler       string      `json:"scheduler"`
	LoraModel       string      `json:"lora_model"`
	Tomesd          string      `json:"tomesd"`
	UseKarrasSigmas string      `json:"use_karras_sigmas"`
	Vae             interface{} `json:"vae"`
	LoraStrength    string      `json:"lora_strength"`
	EmbeddingsModel string      `json:"embeddings_model"`
	Seed            *int64      `json:"seed"`
	Webhook         string      `json:"webhook"`
	TrackId         string      `json:"track_id"`
}

type InPaintMeta struct {
	Prompt         string  `json:"prompt"`
	ModelId        string  `json:"model_id"`
	Scheduler      string  `json:"scheduler"`
	Safetychecker  string  `json:"safetychecker"`
	NegativePrompt string  `json:"negative_prompt"`
	W              int     `json:"W"`
	H              int     `json:"H"`
	GuidanceScale  float64 `json:"guidance_scale"`
	InitImage      string  `json:"init_image"`
	MaskImage      string  `json:"mask_image"`
	MultiLingual   string  `json:"multi_lingual"`
	Steps          int     `json:"steps"`
	NSamples       int     `json:"n_samples"`
	FullUrl        string  `json:"full_url"`
	Upscale        string  `json:"upscale"`
	Seed           int     `json:"seed"`
	Outdir         string  `json:"outdir"`
	FilePrefix     string  `json:"file_prefix"`
}

type InPaintResponse struct {
	Status         string      `json:"status"`
	GenerationTime float64     `json:"generationTime"`
	Id             int         `json:"id"`
	Output         []string    `json:"output"`
	Meta           InPaintMeta `json:"meta"`
}

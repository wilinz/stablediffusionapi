package model

type Text2ImgRequest struct {
	Key               string      `json:"key"`
	ModelId           string      `json:"model_id"`
	Prompt            string      `json:"prompt"`
	NegativePrompt    string      `json:"negative_prompt"`
	Width             string      `json:"width"`
	Height            string      `json:"height"`
	Samples           string      `json:"samples"`
	NumInferenceSteps string      `json:"num_inference_steps"`
	SafetyChecker     string      `json:"safety_checker"`
	EnhancePrompt     string      `json:"enhance_prompt"`
	Seed              *int64      `json:"seed"`
	GuidanceScale     float64     `json:"guidance_scale"`
	MultiLingual      string      `json:"multi_lingual"`
	Panorama          string      `json:"panorama"`
	SelfAttention     string      `json:"self_attention"`
	Upscale           string      `json:"upscale"`
	EmbeddingsModel   string      `json:"embeddings_model"`
	LoraModel         string      `json:"lora_model"`
	Tomesd            string      `json:"tomesd"`
	ClipSkip          string      `json:"clip_skip"`
	UseKarrasSigmas   string      `json:"use_karras_sigmas"`
	Vae               interface{} `json:"vae"`
	LoraStrength      string      `json:"lora_strength"`
	Scheduler         string      `json:"scheduler"`
	Webhook           string      `json:"webhook"`
	TrackId           string      `json:"track_id"`
}

type Text2ImgResponse struct {
	Status         string       `json:"status"`
	GenerationTime float64      `json:"generationTime"`
	Id             int          `json:"id"`
	Output         []string     `json:"output"`
	Meta           Text2ImgMeta `json:"meta"`
}

type Text2ImgMeta struct {
	Prompt         string      `json:"prompt"`
	ModelId        string      `json:"model_id"`
	NegativePrompt string      `json:"negative_prompt"`
	Scheduler      string      `json:"scheduler"`
	Safetychecker  string      `json:"safetychecker"`
	W              int         `json:"W"`
	H              int         `json:"H"`
	GuidanceScale  float64     `json:"guidance_scale"`
	Seed           int64       `json:"seed"`
	Steps          int         `json:"steps"`
	NSamples       int         `json:"n_samples"`
	FullUrl        string      `json:"full_url"`
	Upscale        string      `json:"upscale"`
	MultiLingual   string      `json:"multi_lingual"`
	Panorama       string      `json:"panorama"`
	SelfAttention  string      `json:"self_attention"`
	Embeddings     interface{} `json:"embeddings"`
	Lora           interface{} `json:"lora"`
	Outdir         string      `json:"outdir"`
	FilePrefix     string      `json:"file_prefix"`
}

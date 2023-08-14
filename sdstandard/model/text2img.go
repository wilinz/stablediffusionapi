package model

type Text2ImgRequest struct {
	Key               string  `json:"key"`
	Prompt            string  `json:"prompt"`
	NegativePrompt    string  `json:"negative_prompt"`
	Width             string  `json:"width"`
	Height            string  `json:"height"`
	Samples           string  `json:"samples"`
	NumInferenceSteps string  `json:"num_inference_steps"`
	Seed              *int64  `json:"seed"`
	GuidanceScale     float64 `json:"guidance_scale"`
	SafetyChecker     string  `json:"safety_checker"`
	MultiLingual      string  `json:"multi_lingual"`
	Panorama          string  `json:"panorama"`
	SelfAttention     string  `json:"self_attention"`
	Upscale           string  `json:"upscale"`
	EmbeddingsModel   string  `json:"embeddings_model"`
	Webhook           string  `json:"webhook"`
	TrackId           string  `json:"track_id"`
}

type Text2ImgResponse struct {
	Status         string           `json:"status"`
	GenerationTime float64          `json:"generationTime"`
	ID             int              `json:"id"`
	Output         []string         `json:"output"`
	Meta           Text2ImgMetaData `json:"meta"`
}

type Text2ImgMetaData struct {
	H                      int     `json:"H"`
	W                      int     `json:"W"`
	EnableAttentionSlicing string  `json:"enable_attention_slicing"`
	FilePrefix             string  `json:"file_prefix"`
	GuidanceScale          float64 `json:"guidance_scale"`
	Model                  string  `json:"model"`
	NSamples               int     `json:"n_samples"`
	NegativePrompt         string  `json:"negative_prompt"`
	OutDir                 string  `json:"outdir"`
	Prompt                 string  `json:"prompt"`
	Revision               string  `json:"revision"`
	SafetyChecker          string  `json:"safetychecker"`
	Seed                   int     `json:"seed"`
	Steps                  int     `json:"steps"`
	VAE                    string  `json:"vae"`
}

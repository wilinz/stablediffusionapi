package model

type Img2ImgRequest struct {
	Key               string  `json:"key"`
	Prompt            string  `json:"prompt"`
	NegativePrompt    string  `json:"negative_prompt"`
	InitImage         string  `json:"init_image"`
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

type Img2ImgResponse struct {
	Status         string          `json:"status"`
	GenerationTime float64         `json:"generationTime"`
	ID             int             `json:"id"`
	Output         []string        `json:"output"`
	Meta           Img2ImgMetaData `json:"meta"`
}

type Img2ImgMetaData struct {
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
	SafetyChecker          string  `json:"safety_checker"`
	Seed                   int     `json:"seed"`
	Steps                  int     `json:"steps"`
	VAE                    string  `json:"vae"`
}

package model

type InteriorRequest struct {
	Key           string `json:"key"`
	InitImage     string `json:"init_image"`
	Prompt        string `json:"prompt"`
	Steps         int    `json:"steps"`
	GuidanceScale int    `json:"guidance_scale"`
}

type InteriorResponse struct {
	Status         string       `json:"status"`
	GenerationTime float64      `json:"generationTime"`
	Id             int          `json:"id"`
	Output         []string     `json:"output"`
	Meta           InteriorMeta `json:"meta"`
}

type InteriorMeta struct {
	Prompt          string      `json:"prompt"`
	ModelId         string      `json:"model_id"`
	ControlnetModel string      `json:"controlnet_model"`
	ControlnetType  string      `json:"controlnet_type"`
	NegativePrompt  string      `json:"negative_prompt"`
	Scheduler       string      `json:"scheduler"`
	Safetychecker   string      `json:"safetychecker"`
	AutoHint        string      `json:"auto_hint"`
	GuessMode       string      `json:"guess_mode"`
	Strength        float64     `json:"strength"`
	W               int         `json:"W"`
	H               int         `json:"H"`
	GuidanceScale   int         `json:"guidance_scale"`
	Seed            int64       `json:"seed"`
	MultiLingual    string      `json:"multi_lingual"`
	InitImage       string      `json:"init_image"`
	MaskImage       interface{} `json:"mask_image"`
	Steps           int         `json:"steps"`
	FullUrl         string      `json:"full_url"`
	Upscale         string      `json:"upscale"`
	NSamples        int         `json:"n_samples"`
	Embeddings      interface{} `json:"embeddings"`
	Lora            interface{} `json:"lora"`
	Outdir          string      `json:"outdir"`
	FilePrefix      string      `json:"file_prefix"`
}

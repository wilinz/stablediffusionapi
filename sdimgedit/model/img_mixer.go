package model

type ImgMixerRequest struct {
	Key              string `json:"key"`
	Seed             int    `json:"seed"`
	InitImage        string `json:"init_image"`
	Prompt           string `json:"prompt"`
	NegativePrompt   string `json:"negative_prompt"`
	InitImageWeights string `json:"init_image_weights"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	GuidanceScale    int    `json:"guidance_scale"`
	Steps            int    `json:"steps"`
	Samples          int    `json:"samples"`
}

type ImgMixerResponse struct {
	Status         string       `json:"status"`
	GenerationTime float64      `json:"generationTime"`
	Id             int          `json:"id"`
	Output         []string     `json:"output"`
	Meta           ImgMixerMeta `json:"meta"`
}

type ImgMixerMeta struct {
	H                int    `json:"H"`
	W                int    `json:"W"`
	FilePrefix       string `json:"file_prefix"`
	GuidanceScale    int    `json:"guidance_scale"`
	InitImage        string `json:"init_image"`
	InitImageWeights string `json:"init_image_weights"`
	NSamples         int    `json:"n_samples"`
	NegativePrompt   string `json:"negative_prompt"`
	Outdir           string `json:"outdir"`
	Prompt           string `json:"prompt"`
	Seed             int    `json:"seed"`
	Steps            int    `json:"steps"`
}

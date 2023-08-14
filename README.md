# github.com/wilinz/stablediffusionapi

# Illustrate
This library is https://stablediffusionapi.com/ Go SDK (unofficial)
API documentation address: https://stablediffusionapi.com/docs/

# How to use

```go
package github.com/wilinz/stablediffusionapi

import (
"fmt"
"os"
"github.com/wilinz/stablediffusionapi/sdcommunity"
"github.com/wilinz/stablediffusionapi/sdconfig"
"github.com/wilinz/stablediffusionapi/sdinit"
"github.com/wilinz/stablediffusionapi/sdstandard"
"github.com/wilinz/stablediffusionapi/sdstandard/model"
"testing"
)

func init() {
	KEY := os.Getenv("SD_KEY")
	sdinit.Init(sdconfig.SDConfig{
		ApiKey: KEY,
	})
}

func TestText2Img(t *testing.T) {
	result, _ := sdstandard.Text2Img(model.Text2ImgRequest{
		Prompt:            "ultra realistic close up portrait ((beautiful pale cyberpunk female with heavy black eyeliner))",
		NegativePrompt:    "",
		Width:             "512",
		Height:            "512",
		Samples:           "1",
		NumInferenceSteps: "20",
		Seed:              nil,
		GuidanceScale:     7.5,
		SafetyChecker:     "no",
		MultiLingual:      "no",
		Panorama:          "no",
		SelfAttention:     "no",
		Upscale:           "no",
		EmbeddingsModel:   "",
		Webhook:           "",
		TrackId:           "",
	})
	fmt.Printf("%#v\n", result)
}

func TestImg2Img(t *testing.T) {
	result, _ := sdstandard.Img2Img(model.Img2ImgRequest{
		Prompt:            "a cat sitting on a bench",
		NegativePrompt:    "",
		InitImage:         "https://raw.githubusercontent.com/CompVis/stable-diffusion/main/data/inpainting_examples/overture-creations-5sI6fQgYIuo.png",
		Width:             "512",
		Height:            "512",
		Samples:           "1",
		NumInferenceSteps: "30",
		SafetyChecker:     "no",
		EnhancePrompt:     "yes",
		GuidanceScale:     7.5,
		Strength:          0.7,
		Seed:              nil,
		Webhook:           "",
		TrackId:           "",
	})
	fmt.Printf("%#v\n", result)
}

func TestInPainting(t *testing.T) {
	result, _ := sdstandard.InPaint(model.InPaintRequest{
		Prompt:            "a cat sitting on a bench",
		NegativePrompt:    "",
		InitImage:         "https://raw.githubusercontent.com/CompVis/stable-diffusion/main/data/inpainting_examples/overture-creations-5sI6fQgYIuo.png",
		MaskImage:         "https://raw.githubusercontent.com/CompVis/stable-diffusion/main/data/inpainting_examples/overture-creations-5sI6fQgYIuo_mask.png",
		Width:             "512",
		Height:            "512",
		Samples:           "1",
		NumInferenceSteps: "30",
		SafetyChecker:     "no",
		EnhancePrompt:     "yes",
		GuidanceScale:     7.5,
		Strength:          0.7,
		Seed:              nil,
		Webhook:           "",
		TrackId:           "",
	})
	fmt.Printf("%#v\n", result)
}

func TestTextToVideo(t *testing.T) {
	result, _ := sdstandard.TextToVideo(model.TextToVideoRequest{
		Prompt:         "man walking on the road, ultra HD video",
		NegativePrompt: "Low Quality",
		Scheduler:      "UniPCMultistepScheduler",
		Seconds:        3,
	})
	fmt.Printf("%#v\n", result)
}

func TestModelList(t *testing.T) {
	result, _ := sdcommunity.ModelList()
	fmt.Printf("%#v\n", result)
}

```
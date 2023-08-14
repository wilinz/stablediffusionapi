package model

type ModelReloadRequest struct {
	Key     string `json:"key"`
	ModelId string `json:"model_id"`
}

package model

type FetchQueuedImagesResponse struct {
	Status string   `json:"status"`
	Id     int      `json:"id"`
	Output []string `json:"output"`
}

type FetchQueuedImagesRequest struct {
	Key       string `json:"key"`
	RequestId string `json:"request_id"`
}

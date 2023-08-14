package model

type FetchQueuedImagesResponse struct {
	Status string   `json:"status"`
	Id     int      `json:"id"`
	Output []string `json:"output"`
}

package model

type SystemLoadResponse struct {
	QueueNum  int    `json:"queue_num"`
	QueueTime int    `json:"queue_time"`
	Status    string `json:"status"`
}

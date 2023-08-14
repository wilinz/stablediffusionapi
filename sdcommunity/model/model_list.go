package model

type Model struct {
	ModelId        string      `json:"model_id"`
	Status         string      `json:"status"`
	CreatedAt      interface{} `json:"created_at"`
	InstancePrompt *string     `json:"instance_prompt"`
	ModelName      string      `json:"model_name"`
	Description    string      `json:"description"`
	Screenshots    string      `json:"screenshots"`
}

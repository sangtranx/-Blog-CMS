package chatmodel

import (
	"net/http"
	"time"
)

const EntityName = "Chat"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DeepSeekRequest struct {
	Model    string       `json:"model" gorm:"model"`
	ApiKey   string       `json:"api_key" gorm:"api_key"`
	ApiURL   string       `json:"api_url" gorm:"api_url"`
	Messages []Message    `json:"messages" gorm:"messages"`
	Client   *http.Client `json:"-" gorm:"-"`
}

type DeepSeekResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewDeepSeekModel(model, apiKey, apiURL string, messages []Message) *DeepSeekRequest {
	return &DeepSeekRequest{
		Model:    model,
		ApiKey:   apiKey,
		ApiURL:   apiURL,
		Messages: messages,
		Client:   &http.Client{Timeout: 10 * time.Second},
	}
}

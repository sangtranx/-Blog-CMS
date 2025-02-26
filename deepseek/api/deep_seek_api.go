package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type DeepSeekRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type DeepSeekResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	apiKey := ""
	url := "https://api.deepseek.com/v1/chat/completions"

	requestBody := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []Message{
			{
				Role:    "user",
				Content: "Hello, How are you?",
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error while encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Status Code:", resp.StatusCode)
	switch resp.StatusCode {
	case 200:
		fmt.Println("Yrequest successful, processing response...")
	case 401:
		fmt.Println("Error: Invalid API key.")
		return
	case 403:
		fmt.Println("Error: No API access.")
		return
	case 429:
		fmt.Println("Error: Request limit exceeded.")
		return
	case 500:
		fmt.Println("Error: API server encountered a problem.")
		return
	default:
		fmt.Println("Error: Unknown status code:", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response :", err)
		return
	}

	fmt.Println("Response from server:", string(body))

	var deepSeekResp DeepSeekResponse
	err = json.Unmarshal(body, &deepSeekResp)
	if err != nil {
		fmt.Println("Error when parsing JSON response:", err)
		return
	}

	if len(deepSeekResp.Choices) > 0 {
		fmt.Println("Answer from DeepSeek:", deepSeekResp.Choices[0].Message.Content)
	} else {
		fmt.Println("No response received from API. Please check status code and response content above.")
	}
}

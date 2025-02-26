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
	// Thông tin API
	apiKey := ""
	url := "https://api.deepseek.com/v1/chat/completions"

	// Tạo request body
	requestBody := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []Message{
			{
				Role:    "user",
				Content: "Xin chào, bạn khỏe không? Hôm nay thời tiết thế nào?",
			},
		},
	}

	// Chuyển request body thành JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Lỗi khi mã hóa JSON:", err)
		return
	}

	// Tạo HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Lỗi khi tạo request:", err)
		return
	}

	// Thêm header cho request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Gửi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Lỗi khi gửi request:", err)
		return
	}
	defer resp.Body.Close()

	// Kiểm tra mã trạng thái HTTP
	fmt.Println("Mã trạng thái HTTP:", resp.StatusCode)
	switch resp.StatusCode {
	case 200:
		fmt.Println("Yêu cầu thành công, đang xử lý phản hồi...")
	case 401:
		fmt.Println("Lỗi: API key không hợp lệ.")
		return
	case 403:
		fmt.Println("Lỗi: Không có quyền truy cập API.")
		return
	case 429:
		fmt.Println("Lỗi: Vượt quá giới hạn yêu cầu.")
		return
	case 500:
		fmt.Println("Lỗi: Server API gặp sự cố.")
		return
	default:
		fmt.Println("Lỗi: Mã trạng thái không xác định:", resp.StatusCode)
	}

	// Đọc response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Lỗi khi đọc response:", err)
		return
	}

	// In nội dung phản hồi từ server
	fmt.Println("Phản hồi từ server:", string(body))

	// Parse response JSON
	var deepSeekResp DeepSeekResponse
	err = json.Unmarshal(body, &deepSeekResp)
	if err != nil {
		fmt.Println("Lỗi khi parse JSON response:", err)
		return
	}

	// Kiểm tra và in câu trả lời
	if len(deepSeekResp.Choices) > 0 {
		fmt.Println("Câu trả lời từ DeepSeek:", deepSeekResp.Choices[0].Message.Content)
	} else {
		fmt.Println("Không nhận được câu trả lời từ API. Vui lòng kiểm tra mã trạng thái và nội dung phản hồi ở trên.")
	}
}

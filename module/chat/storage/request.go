package chatstorage

import (
	chatmodel "Blog-CMS/module/chat/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *sqlStorage) CallDeepSeekAPI(reqData chatmodel.DeepSeekRequest) (*chatmodel.DeepSeekResponse, error) {
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", reqData.ApiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+reqData.ApiKey)

	resp, err := reqData.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	var deepSeekResp chatmodel.DeepSeekResponse
	if err := json.Unmarshal(body, &deepSeekResp); err != nil {
		return nil, err
	}

	return &deepSeekResp, nil
}

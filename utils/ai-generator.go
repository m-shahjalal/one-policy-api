package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/m-shahjalal/onepolicy-api/config"
)

type APIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func GeneratePolicy(data any) (string, error) {
	payload := map[string]any{
		"model":       "deepseek/deepseek-chat-v3-0324:free",
		"temperature": 0.3,
		"max_tokens":  4000,
		"messages": []map[string]any{{
			"role":    "user",
			"content": buildPrompt(data),
		}},
	}

	body, err := makeRequest("https://openrouter.ai/api/v1/chat/completions", payload)
	if err != nil {
		return "", err
	}

	var resp APIResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", fmt.Errorf("API error: %s", resp.Error.Message)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices")
	}

	return resp.Choices[0].Message.Content, nil
}

func buildPrompt(data any) string {
	var dataStr string

	switch v := data.(type) {
	case string:
		dataStr = v
	case []byte:
		dataStr = string(v)
	default:
		if jsonBytes, err := json.MarshalIndent(v, "", "  "); err == nil {
			dataStr = string(jsonBytes)
		} else {
			dataStr = fmt.Sprintf("%+v", v)
		}
	}

	return fmt.Sprintf(config.PromptDirectives, dataStr)
}

func makeRequest(url string, data any) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENROUTER_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OllamaClient struct {
	baseURL string
	model   string
	client  *http.Client
}

func NewOllamaClient(baseURL string, model string) *OllamaClient {
	return &OllamaClient{
		baseURL: baseURL,
		model:   model,
		client: &http.Client{
			Timeout: 2 * time.Minute,
		},
	}
}

func (o *OllamaClient) Chat(ctx context.Context, messages []Message) (string, error) {

	reqBody := ChatRequest{
		Model:    o.model,
		Messages: messages,
		Stream:   false,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		o.baseURL+"/api/chat",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := o.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"ollama returned status %d",
			resp.StatusCode,
		)
	}

	var response ChatResponse

	err = json.NewDecoder(resp.Body).
		Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Message.Content, nil
}

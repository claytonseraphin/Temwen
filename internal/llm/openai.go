package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/claytonseraphin/temwen/internal/config"
)

type OpenAIClient struct {
	model config.Model
	http  *http.Client
}

type chatResponse struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Content *string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
}

func NewOpenAIClient(m config.Model) *OpenAIClient {
	return &OpenAIClient{model: m, http: &http.Client{Timeout: m.Timeout.Duration}}
}

func (c *OpenAIClient) Chat(ctx context.Context, messages []Message) (Reply, error) {
	body := map[string]any{
		"model":       c.model.Model,
		"messages":    messages,
		"max_tokens":  c.model.MaxTokens,
		"temperature": c.model.Temperature,
	}
	for k, v := range c.model.Extra {
		body[k] = v
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return Reply{}, fmt.Errorf("error encoding payload, %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.model.URL, bytes.NewReader(payload))
	if err != nil {
		return Reply{}, fmt.Errorf("error while processing request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()

	resp, err := c.http.Do(req)
	if err != nil {
		return Reply{}, fmt.Errorf("error while processing response: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return Reply{}, fmt.Errorf("status %d: %s", resp.StatusCode, msg)
	}

	var out chatResponse

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return Reply{}, fmt.Errorf("decode response: %w", err)
	}

	if len(out.Choices) == 0 {
		return Reply{}, fmt.Errorf("response has no choices")
	}
	choice := out.Choices[0]
	if choice.Message.Content == nil {
		return Reply{}, fmt.Errorf("model returned no content (finish_reason=%s); is thinking mode on?", choice.FinishReason)
	}

	return Reply{
		Content:      *choice.Message.Content,
		FinishReason: choice.FinishReason,
		InputTokens:  out.Usage.PromptTokens,
		OutputTokens: out.Usage.CompletionTokens,
		Latency:      time.Since(start),
	}, nil
}

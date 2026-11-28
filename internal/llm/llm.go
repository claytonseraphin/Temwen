package llm

import (
	"context"
	"time"
)

type Message struct {
	Role    string `json:"role"` //system user, or assistant
	Content string `json:"content"`
}

type Reply struct {
	Content      string
	FinishReason string
	InputTokens  int
	OutputTokens int
	Latency      time.Duration
}

type ChatModel interface {
	Chat(ctx context.Context, messages []Message) (Reply, error)
}

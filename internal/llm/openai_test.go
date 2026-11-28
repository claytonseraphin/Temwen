package llm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/claytonseraphin/temwen/internal/config"
)

func TestOpenAIClientChat(t *testing.T) {
	// 1. A fake server that records what the client sent and returns a canned reply.
	var got map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Errorf("server could not decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"choices": [{"finish_reason": "stop", "message": {"content": "Bonjou!"}}],
			"usage": {"prompt_tokens": 12, "completion_tokens": 3}
		}`))
	}))
	defer server.Close()

	// 2. The real client, pointed at the fake server.
	client := NewOpenAIClient(config.Model{
		URL:       server.URL,
		Model:     "test-model",
		MaxTokens: 50,
		Extra: map[string]any{
			"chat_template_kwargs": map[string]any{"enable_thinking": false},
		},
	})

	messages := []Message{
		{Role: "system", Content: "test system prompt"},
		{Role: "user", Content: "Hello"},
	}

	reply, err := client.Chat(context.Background(), messages)
	if err != nil {
		t.Fatalf("Chat returned an error: %v", err)
	}

	// 3. Check the reply was read correctly.
	if reply.Content != "Bonjou!" {
		t.Errorf("Content = %q, want %q", reply.Content, "Bonjou!")
	}
	if reply.FinishReason != "stop" {
		t.Errorf("FinishReason = %q, want %q", reply.FinishReason, "stop")
	}
	if reply.InputTokens != 12 || reply.OutputTokens != 3 {
		t.Errorf("tokens = %d in / %d out, want 12 / 3", reply.InputTokens, reply.OutputTokens)
	}

	// 4. Check what the client sent.
	if got["model"] != "test-model" {
		t.Errorf("sent model = %v, want test-model", got["model"])
	}
	if got["max_tokens"] != float64(50) { // JSON numbers decode as float64
		t.Errorf("sent max_tokens = %v, want 50", got["max_tokens"])
	}
	if _, ok := got["chat_template_kwargs"]; !ok {
		t.Error("chat_template_kwargs was not sent at the top level of the body")
	}
	sent, _ := got["messages"].([]any)
	if len(sent) != 2 {
		t.Fatalf("sent %d messages, want 2", len(sent))
	}
	first, _ := sent[0].(map[string]any)
	if first["role"] != "system" {
		t.Errorf("first message role = %v, want system (are the json tags lowercase?)", first["role"])
	}
}

func TestOpenAIClientErrors(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		body    string
		wantErr string
	}{
		{"server error", 500, `model not found`, "status 500"},
		{"null content (thinking mode)", 200, `{"choices": [{"finish_reason": "length", "message": {"content": null}}]}`, "no content"},
		{"no choices", 200, `{"choices": []}`, "no choices"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				w.Write([]byte(tt.body))
			}))
			defer server.Close()

			client := NewOpenAIClient(config.Model{URL: server.URL, MaxTokens: 50})
			_, err := client.Chat(context.Background(), []Message{{Role: "user", Content: "Hello"}})
			if err == nil {
				t.Fatalf("expected an error containing %q, got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("expected an error containing %q, got %q", tt.wantErr, err)
			}
		})
	}
}

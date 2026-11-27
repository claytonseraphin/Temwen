package config

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	path := filepath.Join("..", "..", "config", "registry.yaml")

	reg, err := Load(path)
	if err != nil {
		t.Fatalf("Load(%q) returned an error: %v", path, err)
	}

	if len(reg.Models) == 0 {
		t.Fatal("expected models in the registry, got none")
	}
	for id, m := range reg.Models {
		if m.ID != id {
			t.Errorf("model %s: ID = %q, want %q", id, m.ID, id)
		}
	}

}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		witnesses map[string]Witness
		wantErr   bool
	}{
		{"valid handoff", map[string]Witness{
			"w1": {HandoffTo: []string{"w2"}},
			"w2": {HandoffTo: []string{"w1"}},
		}, false},
		{"unknown target", map[string]Witness{
			"w1": {HandoffTo: []string{"w2"}},
		}, true},
		{"hands off to itself", map[string]Witness{
			"w1": {HandoffTo: []string{"w1"}},
		}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := Registry{Witnesses: tt.witnesses}
			err := reg.validate()
			if (err != nil) != tt.wantErr {
				t.Fatalf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestValidateModels(t *testing.T) {
	// A model that passes every check; each case changes one thing.
	qwen := Model{Provider: "openai-compatible", URL: "http://localhost:8080/v1/chat/completions", MaxTokens: 250}

	tests := []struct {
		name         string
		defaultModel string
		models       map[string]Model
		wantErr      string // empty means no error expected
	}{
		{
			name:         "valid config",
			defaultModel: "qwen-local",
			models: map[string]Model{
				"qwen-local":   qwen,
				"claude-haiku": {Provider: "anthropic", MaxTokens: 250},
			},
		},
		{
			name:         "default model missing",
			defaultModel: "qwen-loca",
			models:       map[string]Model{"qwen-local": qwen},
			wantErr:      "is not a defined model",
		},
		{
			name:         "default model empty",
			defaultModel: "",
			models:       map[string]Model{"qwen-local": qwen},
			wantErr:      "is not a defined model",
		},
		{
			name:         "unknown provider",
			defaultModel: "qwen-local",
			models: map[string]Model{
				"qwen-local": {Provider: "openia", URL: qwen.URL, MaxTokens: 250},
			},
			wantErr: "unknown provider",
		},
		{
			name:         "openai-compatible without url",
			defaultModel: "qwen-local",
			models: map[string]Model{
				"qwen-local": {Provider: "openai-compatible", MaxTokens: 250},
			},
			wantErr: "url is required",
		},
		{
			name:         "anthropic without url is fine",
			defaultModel: "claude-haiku",
			models: map[string]Model{
				"claude-haiku": {Provider: "anthropic", MaxTokens: 250},
			},
		},
		{
			name:         "max_tokens zero",
			defaultModel: "qwen-local",
			models: map[string]Model{
				"qwen-local": {Provider: qwen.Provider, URL: qwen.URL, MaxTokens: 0},
			},
			wantErr: "max_tokens must be greater than 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := Registry{DefaultModel: tt.defaultModel, Models: tt.models}
			err := reg.validateModels()

			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected an error containing %q, got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("expected an error containing %q, got %q", tt.wantErr, err)
			}
		})
	}
}

package config

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.yaml.in/yaml/v3"
)

type Date struct{ time.Time }
type Duration struct{ time.Duration }

// helper function to convert registry era date
func (d *Date) UnmarshalYAML(n *yaml.Node) error {
	t, err := time.Parse("2006-01-02", n.Value)
	if err != nil {
		return fmt.Errorf("invalid date %q, want YYYY-MM-DD: %w", n.Value, err)
	}
	d.Time = t
	return nil
}

func (d *Duration) UnmarshalYAML(n *yaml.Node) error {
	v, err := time.ParseDuration(n.Value)
	if err != nil {
		return fmt.Errorf("invalid duration %q, want e.g. 60s or 2m: %w", n.Value, err)
	}
	d.Duration = v
	return nil
}

type Registry struct {
	Witnesses    map[string]Witness `yaml:"witnesses"`
	DefaultModel string             `yaml:"default_model"`
	Models       map[string]Model   `yaml:"models"`
}

type Era struct {
	Start Date `yaml:"start"`
	End   Date `yaml:"end"`
}

type Witness struct {
	ID            string   `yaml:"-"`
	Name          string   `yaml:"name"`
	Era           Era      `yaml:"era"`
	Prompt        string   `yaml:"prompt"`
	SourceShelf   string   `yaml:"source_shelf"`
	HandoffTo     []string `yaml:"handoff_to"`
	Voice         string   `yaml:"voice"`
	PromptText    string   `yaml:"-"`
	PromptVersion string   `yaml:"-"`
}

type Model struct {
	ID          string         `yaml:"id"`
	Provider    string         `yaml:"provider"`
	URL         string         `yaml:"url"`
	Model       string         `yaml:"model"`
	APIKeyEnv   string         `yaml:"api_key_env"`
	MaxTokens   int            `yaml:"max_tokens"`
	Temperature float32        `yaml:"temperature"`
	Timeout     Duration       `yaml:"timeout"`
	Extra       map[string]any `yaml:"extra"`
}

func ReadPrompt(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read prompt %s: %w", path, err)
	}
	return string(data), nil
}

func ComputePromptVersion(input string) string {
	hashBytes := sha256.Sum256([]byte(input))

	return hex.EncodeToString(hashBytes[:])
}

func (reg *Registry) validate() error {
	for id, w := range reg.Witnesses {
		for _, target := range w.HandoffTo {
			if target == id {
				return fmt.Errorf("witness %s: cannot hand off to itself", id)
			}
			if _, ok := reg.Witnesses[target]; !ok {
				return fmt.Errorf("witness %s: handoff_to %q is not a known witness", id, target)
			}
		}
	}

	return nil
}

var knownProviders = map[string]bool{
	"openai-compatible": true,
	"anthropic":         true,
}

func (reg *Registry) validateModels() error {
	if _, ok := reg.Models[reg.DefaultModel]; !ok {
		return fmt.Errorf("default_model %q is not a defined model", reg.DefaultModel)
	}

	for id, m := range reg.Models {
		if !knownProviders[m.Provider] {
			return fmt.Errorf("model %s: unknown provider %q", id, m.Provider)
		}

		if m.Provider == "openai-compatible" && m.URL == "" {
			return fmt.Errorf("model %s: url is required for provider %q", id, m.Provider)
		}

		if m.MaxTokens <= 0 {
			return fmt.Errorf("model %s: max_tokens must be greater than 0, got %d", id, m.MaxTokens)
		}
	}

	return nil
}

func Load(path string) (*Registry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read registry %s: %w", path, err)
	}

	var reg Registry
	if err := yaml.Unmarshal(data, &reg); err != nil {
		return nil, fmt.Errorf("decode registry %s: %w", path, err)
	}

	base := filepath.Dir(path)

	for id, w := range reg.Witnesses {
		w.ID = id

		text, err := ReadPrompt(filepath.Join(base, w.Prompt))
		if err != nil {
			return nil, fmt.Errorf("witness %s: %w", id, err)
		}
		w.PromptText = text

		version := ComputePromptVersion(w.PromptText)
		w.PromptVersion = version

		reg.Witnesses[id] = w
	}
	fmt.Printf("Registry: %v", reg.DefaultModel)

	for id, m := range reg.Models {
		m.ID = id
		reg.Models[id] = m
	}

	if err := reg.validate(); err != nil {
		return nil, err
	}

	if err := reg.validateModels(); err != nil {
		return nil, err
	}

	return &reg, nil
}

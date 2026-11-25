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

// helper function to convert registry era date
func (d *Date) UnmarshalYAML(n *yaml.Node) error {
	t, err := time.Parse("2006-01-02", n.Value)
	if err != nil {
		return fmt.Errorf("invalid date %q, want YYYY-MM-DD: %w", n.Value, err)
	}
	d.Time = t
	return nil
}

type Registry struct {
	Witnesses map[string]Witness `yaml:"witnesses"`
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
	for key := range reg.Witnesses {
		fmt.Println("witness key:", key)
	}

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

	if err := reg.validate(); err != nil {
		return nil, err
	}

	return &reg, nil
}

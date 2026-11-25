package config

import (
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	path := filepath.Join("..", "..", "config", "witnesses.yaml")

	_, err := Load(path)
	if err != nil {
		t.Fatalf("Load(%q) returned an error: %v", path, err)
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

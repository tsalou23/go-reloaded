package tests

import (
	"go-reloaded/internal/processor"
	"testing"
)

func TestTrickyCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"C1", "a honest man", "an honest man"},
		{"C2", "10 (bin) and 1A (hex)", "2 and 26"},
		{"C3", "HELLO THERE (low, 2) WORLD", "hello there WORLD"},
		{"C4", "I waited ... and then ?!", "I waited... and then?!"},
		{"C5", "He said ' hello there '", "He said 'hello there'"},
	}

	pipeline := processor.NewPipeline()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.Process(tt.input)
			if result != tt.expected {
				t.Errorf("Test %s failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.name, tt.input, tt.expected, result)
			}
		})
	}
}
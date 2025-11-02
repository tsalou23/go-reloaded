package tests

import (
	"go-reloaded/internal/processor"
	"testing"
)

func TestGoldenCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"T1", "1E (hex) files were added", "30 files were added"},
		{"T2", "It has been 10 (bin) years", "It has been 2 years"},
		{"T3", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"T4", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"T5", "Welcome to the brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"T6", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"T7", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"T8", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"T9", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"T10", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
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
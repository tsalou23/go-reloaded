package tests

import (
	"fmt"
	"go-reloaded/internal/processor"
	"testing"
)

func TestPipelineMode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hex conversion", "1E (hex) files", "30 files"},
		{"Binary conversion", "10 (bin) years", "2 years"},
		{"Uppercase", "go (up)", "GO"},
		{"Lowercase", "LOUD (low)", "loud"},
		{"Capitalize", "bridge (cap)", "Bridge"},
		{"Multi-word up", "so exciting (up, 2)", "SO EXCITING"},
		{"Article correction", "a honest man", "an honest man"},
		{"Quote cleaning", "' hello '", "'hello'"},
		{"Punctuation fix", "Hi , world !", "Hi, world!"},
		{"Complex case", "1A (hex) items (up) and ' test '", "26 ITEMS and 'test'"},
	}

	pipeline := processor.NewPipeline()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.Process(tt.input)
			if result != tt.expected {
				t.Errorf("Pipeline mode failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestFSMMode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hex conversion", "1E (hex) files", "30 files"},
		{"Binary conversion", "10 (bin) years", "2 years"},
		{"Uppercase", "go (up)", "GO"},
		{"Lowercase", "LOUD (low)", "loud"},
		{"Capitalize", "bridge (cap)", "Bridge"},
		{"Multi-word up", "so exciting (up, 2)", "SO EXCITING"},
		{"Article correction", "a honest man", "an honest man"},
		{"Quote cleaning", "' hello '", "'hello'"},
		{"Punctuation fix", "Hi , world !", "Hi, world!"},
		{"Complex case", "1A (hex) items (up) and ' test '", "26 ITEMS and 'test'"},
	}

	fsm := processor.NewFSM()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fsm.Process(tt.input)
			if result != tt.expected {
				t.Errorf("FSM mode failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestHybridMode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Hex conversion", "1E (hex) files", "30 files"},
		{"Binary conversion", "10 (bin) years", "2 years"},
		{"Uppercase", "go (up)", "GO"},
		{"Lowercase", "LOUD (low)", "loud"},
		{"Capitalize", "bridge (cap)", "Bridge"},
		{"Multi-word up", "so exciting (up, 2)", "SO EXCITING"},
		{"Article correction", "a honest man", "an honest man"},
		{"Quote cleaning", "' hello '", "'hello'"},
		{"Punctuation fix", "Hi , world !", "Hi, world!"},
		{"Complex case", "1A (hex) items (up) and ' test '", "26 ITEMS and 'test'"},
	}

	hybrid := processor.NewHybrid()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hybrid.Process(tt.input)
			if result != tt.expected {
				t.Errorf("Hybrid mode failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestModeConsistency(t *testing.T) {
	testCases := []string{
		"1E (hex) files were added",
		"It has been 10 (bin) years",
		"Ready, set, go (up) !",
		"I should stop SHOUTING (low)",
		"Welcome to the brooklyn bridge (cap)",
		"This is so exciting (up, 2)",
		"There it was. A amazing rock!",
		"I am exactly how they describe me: ' awesome '",
	}

	pipeline := processor.NewPipeline()
	fsm := processor.NewFSM()
	hybrid := processor.NewHybrid()

	for i, input := range testCases {
		t.Run(fmt.Sprintf("Consistency_%d", i+1), func(t *testing.T) {
			pipelineResult := pipeline.Process(input)
			fsmResult := fsm.Process(input)
			hybridResult := hybrid.Process(input)

			if pipelineResult != fsmResult {
				t.Errorf("Pipeline and FSM results differ:\nInput: %q\nPipeline: %q\nFSM: %q", input, pipelineResult, fsmResult)
			}

			if pipelineResult != hybridResult {
				t.Errorf("Pipeline and Hybrid results differ:\nInput: %q\nPipeline: %q\nHybrid: %q", input, pipelineResult, hybridResult)
			}
		})
	}
}
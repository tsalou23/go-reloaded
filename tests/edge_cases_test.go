package tests

import (
	"go-reloaded/internal/processor"
	"testing"
)

func TestEdgeCasesPipeline(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Only spaces", "   ", "   "},
		{"Invalid hex", "ZZ (hex)", "ZZ (hex)"},
		{"Invalid binary", "22 (bin)", "22 (bin)"},
		{"Zero count", "word (up, 0)", "word (up, 0)"},
		{"Negative count", "word (up, -1)", "word (up, -1)"},
		{"Large count", "a b c (up, 10)", "A B C (up, 10)"},
		{"Nested quotes", "' hello ' world ' test '", "'hello' world 'test'"},
		{"Multiple punctuation", "Hi !! ?? ..", "Hi!! ?? .."},
		{"Article edge cases", "a hour", "an hour"},
		{"Silent h words", "a honest", "an honest"},
		{"Non-silent h", "a house", "a house"},
	}

	pipeline := processor.NewPipeline()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.Process(tt.input)
			if result != tt.expected {
				t.Errorf("Pipeline edge case failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestEdgeCasesFSM(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Only spaces", "   ", "   "},
		{"Invalid hex", "ZZ (hex)", "ZZ (hex)"},
		{"Invalid binary", "22 (bin)", "22 (bin)"},
		{"Zero count", "word (up, 0)", "word (up, 0)"},
		{"Negative count", "word (up, -1)", "word (up, -1)"},
		{"Large count", "a b c (up, 10)", "A B C (up, 10)"},
		{"Nested quotes", "' hello ' world ' test '", "'hello' world 'test'"},
		{"Multiple punctuation", "Hi !! ?? ..", "Hi!! ?? .."},
		{"Article edge cases", "a hour", "an hour"},
		{"Silent h words", "a honest", "an honest"},
		{"Non-silent h", "a house", "a house"},
	}

	fsm := processor.NewFSM()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fsm.Process(tt.input)
			if result != tt.expected {
				t.Errorf("FSM edge case failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestEdgeCasesHybrid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Only spaces", "   ", "   "},
		{"Invalid hex", "ZZ (hex)", "ZZ (hex)"},
		{"Invalid binary", "22 (bin)", "22 (bin)"},
		{"Zero count", "word (up, 0)", "word (up, 0)"},
		{"Negative count", "word (up, -1)", "word (up, -1)"},
		{"Large count", "a b c (up, 10)", "A B C (up, 10)"},
		{"Nested quotes", "' hello ' world ' test '", "'hello' world 'test'"},
		{"Multiple punctuation", "Hi !! ?? ..", "Hi!! ?? .."},
		{"Article edge cases", "a hour", "an hour"},
		{"Silent h words", "a honest", "an honest"},
		{"Non-silent h", "a house", "a house"},
	}

	hybrid := processor.NewHybrid()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hybrid.Process(tt.input)
			if result != tt.expected {
				t.Errorf("Hybrid edge case failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestPerformanceComparison(t *testing.T) {
	longText := `This is a very long text with 1E (hex) numbers and 10 (bin) values. 
	It contains UPPERCASE (low) words and lowercase (up) words. 
	There are also ' quoted strings ' and punctuation marks , ! ? . 
	We need to test a honest approach to see a amazing performance difference. 
	The text should be processed quickly (cap, 3) by all modes.`

	pipeline := processor.NewPipeline()
	fsm := processor.NewFSM()
	hybrid := processor.NewHybrid()

	// Just verify all modes produce the same result for long text
	pipelineResult := pipeline.Process(longText)
	fsmResult := fsm.Process(longText)
	hybridResult := hybrid.Process(longText)

	if pipelineResult != fsmResult {
		t.Errorf("Pipeline and FSM results differ for long text")
	}

	if pipelineResult != hybridResult {
		t.Errorf("Pipeline and Hybrid results differ for long text")
	}
}
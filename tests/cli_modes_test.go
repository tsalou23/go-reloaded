package tests

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLIModes(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "go-reloaded", "../cmd/go-reloaded")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("go-reloaded")

	testInput := "1E (hex) files and go (up) test"
	expectedOutput := "30 files and GO test"

	modes := []string{"pipeline", "fsm", "hybrid"}

	for _, mode := range modes {
		t.Run("Mode_"+mode, func(t *testing.T) {
			// Create input file
			inputFile := "test-outputs/test_input_" + mode + ".txt"
			outputFile := "test-outputs/test_output_" + mode + ".txt"
			
			err := os.WriteFile(inputFile, []byte(testInput), 0644)
			if err != nil {
				t.Fatalf("Failed to create input file: %v", err)
			}
			defer os.Remove(inputFile)
			defer os.Remove(outputFile)

			// Run the CLI with the mode
			cmd := exec.Command("./go-reloaded", inputFile, outputFile, mode)
			output, err := cmd.CombinedOutput()
			
			if err != nil {
				t.Fatalf("CLI failed for mode %s: %v\nOutput: %s", mode, err, string(output))
			}

			// Check output file exists and has correct content
			content, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatalf("Failed to read output file for mode %s: %v", mode, err)
			}

			result := strings.TrimSpace(string(content))
			if result != expectedOutput {
				t.Errorf("Mode %s failed:\nExpected: %q\nGot:      %q", mode, expectedOutput, result)
			}
		})
	}
}

func TestCLIComplexCases(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "go-reloaded", "../cmd/go-reloaded")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("go-reloaded")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Complex_case_1",
			"1A (hex) items (up) and ' hello world '",
			"26 ITEMS and 'hello world'",
		},
		{
			"Complex_case_2", 
			"There was a honest man , who said ' I am happy (cap) ' !",
			"There was an honest man, who said 'I Am Happy'!",
		},
		{
			"Complex_case_3",
			"The result is 1010 (bin) and EVERYTHING (low, 2) works fine .",
			"The result is 10 and everything works fine.",
		},
	}

	modes := []string{"pipeline", "fsm", "hybrid"}

	for _, mode := range modes {
		for _, tt := range tests {
			t.Run(mode+"_"+tt.name, func(t *testing.T) {
				inputFile := "test-outputs/test_" + mode + "_" + tt.name + "_input.txt"
				outputFile := "test-outputs/test_" + mode + "_" + tt.name + "_output.txt"
				
				err := os.WriteFile(inputFile, []byte(tt.input), 0644)
				if err != nil {
					t.Fatalf("Failed to create input file: %v", err)
				}
				defer os.Remove(inputFile)
				defer os.Remove(outputFile)

				// Run the CLI
				cmd := exec.Command("./go-reloaded", inputFile, outputFile, mode)
				output, err := cmd.CombinedOutput()
				
				if err != nil {
					t.Fatalf("CLI failed for %s %s: %v\nOutput: %s", mode, tt.name, err, string(output))
				}

				// Check result
				content, err := os.ReadFile(outputFile)
				if err != nil {
					t.Fatalf("Failed to read output file: %v", err)
				}

				result := strings.TrimSpace(string(content))
				if result != tt.expected {
					t.Errorf("%s %s failed:\nInput:    %q\nExpected: %q\nGot:      %q", mode, tt.name, tt.input, tt.expected, result)
				}
			})
		}
	}
}
package tests

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "go-reloaded", "../cmd/go-reloaded")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("go-reloaded")

	tests := []struct {
		name     string
		args     []string
		wantExit int
		wantOut  string
	}{
		{
			name:     "Missing arguments",
			args:     []string{},
			wantExit: 1,
			wantOut:  "Usage:",
		},
		{
			name:     "Invalid mode",
			args:     []string{"input.txt", "output.txt", "invalid"},
			wantExit: 1,
			wantOut:  "Error: invalid mode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("./go-reloaded", tt.args...)
			output, err := cmd.CombinedOutput()
			
			exitCode := 0
			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					exitCode = exitError.ExitCode()
				}
			}

			if exitCode != tt.wantExit {
				t.Errorf("Expected exit code %d, got %d", tt.wantExit, exitCode)
			}

			if !strings.Contains(string(output), tt.wantOut) {
				t.Errorf("Expected output to contain %q, got %q", tt.wantOut, string(output))
			}
		})
	}
}
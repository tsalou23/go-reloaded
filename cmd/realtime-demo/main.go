package main

import (
	"bufio"
	"fmt"
	"go-reloaded/internal/processor"
	"go-reloaded/internal/rules"
	"os"
	"strings"
)

func main() {
	fmt.Println("Real-time FSM Demo")
	fmt.Println("Type text with transformations like: hello(up) or 1E(hex)")
	fmt.Println("Press Enter to process, 'quit' to exit")
	fmt.Println()

	fsm := processor.NewRealtimeFSM()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input: ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if strings.ToLower(input) == "quit" {
			break
		}

		// Process character by character
		var output strings.Builder
		for _, char := range input {
			result := fsm.ProcessChar(char)
			if result != "" {
				output.WriteString(result)
			}
		}

		// Show current buffer (incomplete input)
		buffer := fsm.GetCurrentBuffer()
		if buffer != "" {
			output.WriteString(buffer)
		}

		// Apply all final rules to complete output
		finalOutput := output.String()
		finalOutput = rules.CleanQuotes(finalOutput)
		finalOutput = rules.FixPunctuation(finalOutput)
		finalOutput = rules.FixArticles(finalOutput)
		fmt.Printf("Output: %s\n\n", finalOutput)
		fsm.Reset() // Reset for next input
	}
}
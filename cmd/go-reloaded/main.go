package main

import (
	"fmt"
	"go-reloaded/internal/processor"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		printUsage()
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	mode := os.Args[3]

	// Validate mode
	var proc processor.Processor
	switch mode {
	case "pipeline":
		proc = processor.NewPipeline()
	case "fsm":
		proc = processor.NewFSM()
	case "hybrid":
		proc = processor.NewHybrid()
	default:
		fmt.Println("Error: invalid mode. Use one of [pipeline|fsm|hybrid].")
		os.Exit(1)
	}

	// Read input file
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(2)
	}

	// Process text
	result := proc.Process(string(input))

	// Write output file
	err = ioutil.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(2)
	}
}

func printUsage() {
	fmt.Println("Usage: go-reloaded <input_file> <output_file> <mode>")
	fmt.Println("Modes:")
	fmt.Println("  pipeline   Sequential modular processor")
	fmt.Println("  fsm        Finite State Machine processor")
	fmt.Println("  hybrid     FSM tokenizer + pipeline rules")
}
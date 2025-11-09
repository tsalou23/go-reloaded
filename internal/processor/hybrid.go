package processor

import (
	"go-reloaded/internal/rules"
	"go-reloaded/internal/tokenizer"
)

// Hybrid implements the Processor interface using FSM tokenizer + pipeline rules
type Hybrid struct{}

// NewHybrid creates a new hybrid processor
func NewHybrid() *Hybrid {
	return &Hybrid{}
}

// Process applies rules using hybrid approach: FSM tokenizer + pipeline rules
func (h *Hybrid) Process(text string) string {
	// Step 1: Use FSM tokenizer to parse and preprocess the text
	tokenizer := tokenizer.NewTokenizer()
	tokens := tokenizer.Tokenize(text)
	
	// Step 2: Apply smart preprocessing based on token analysis
	preprocessedText := tokenizer.PreprocessTokens(tokens)
	
	// Step 3: Apply pipeline rules to the preprocessed text
	result := rules.ApplyCase(preprocessedText)
	result = rules.ApplyNumbers(result)
	result = rules.CleanQuotes(result)
	result = rules.FixPunctuation(result)
	result = rules.FixArticles(result) // Apply articles last to avoid conflicts
	
	return result
}
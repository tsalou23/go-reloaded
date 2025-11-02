package processor

import "go-reloaded/internal/rules"

// Hybrid implements the Processor interface using FSM tokenizer + pipeline rules
type Hybrid struct{}

// NewHybrid creates a new hybrid processor
func NewHybrid() *Hybrid {
	return &Hybrid{}
}

// Process applies rules using hybrid approach
func (h *Hybrid) Process(text string) string {
	// For now, use same rule order as pipeline
	// TODO: Implement FSM tokenizer + pipeline rule application
	text = rules.ApplyNumbers(text)
	text = rules.ApplyCase(text)
	text = rules.FixArticles(text)
	text = rules.CleanQuotes(text)
	text = rules.FixPunctuation(text)
	return text
}
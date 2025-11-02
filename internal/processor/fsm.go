package processor

import "go-reloaded/internal/rules"

// FSM implements the Processor interface using finite state machine
type FSM struct{}

// NewFSM creates a new FSM processor
func NewFSM() *FSM {
	return &FSM{}
}

// Process applies rules using FSM approach (simplified implementation)
func (f *FSM) Process(text string) string {
	// For now, use same rule order as pipeline
	// TODO: Implement proper FSM tokenization and state transitions
	text = rules.ApplyNumbers(text)
	text = rules.ApplyCase(text)
	text = rules.FixArticles(text)
	text = rules.CleanQuotes(text)
	text = rules.FixPunctuation(text)
	return text
}
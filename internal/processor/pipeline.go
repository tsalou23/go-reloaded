package processor

import "go-reloaded/internal/rules"

// Pipeline implements the Processor interface using sequential rule application
type Pipeline struct{}

// NewPipeline creates a new pipeline processor
func NewPipeline() *Pipeline {
	return &Pipeline{}
}

// Process applies all rules in the specified order
func (p *Pipeline) Process(text string) string {
	text = rules.ApplyCase(text)
	text = rules.ApplyNumbers(text)
	text = rules.CleanQuotes(text)
	text = rules.FixPunctuation(text)
	text = rules.FixArticles(text) // Apply articles last to avoid conflicts
	return text
}
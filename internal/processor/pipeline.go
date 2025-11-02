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
	text = rules.ApplyNumbers(text)
	text = rules.ApplyCase(text)
	text = rules.FixArticles(text)
	text = rules.CleanQuotes(text)
	text = rules.FixPunctuation(text)
	return text
}
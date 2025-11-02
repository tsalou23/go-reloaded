package processor

// Processor defines the interface for text processing
type Processor interface {
	Process(text string) string
}
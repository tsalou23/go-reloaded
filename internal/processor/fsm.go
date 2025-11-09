package processor

import (
	"go-reloaded/internal/rules"
	"strconv"
	"strings"
)

// FSMState represents the current state of the FSM
type FSMState int

const (
	Normal FSMState = iota
	InMarker
	InQuotes
)

// FSM implements the Processor interface using finite state machine
type FSM struct {
	state FSMState
}

// NewFSM creates a new FSM processor
func NewFSM() *FSM {
	return &FSM{state: Normal}
}

// Process applies rules using FSM approach with character-by-character processing
func (f *FSM) Process(text string) string {
	// FSM processes text character by character, tracking state
	result := f.processWithFSM(text)
	
	// Apply articles last to avoid conflicts
	result = rules.FixArticles(result)
	return result
}

// processWithFSM uses finite state machine to process text character by character
func (f *FSM) processWithFSM(text string) string {
	var result strings.Builder
	var currentWord strings.Builder
	var markerContent strings.Builder
	
	f.state = Normal
	runes := []rune(text)
	
	for _, char := range runes {
		switch f.state {
		case Normal:
			if char == '(' {
				// Save current word and enter marker state
				if currentWord.Len() > 0 {
					result.WriteString(currentWord.String())
					currentWord.Reset()
				}
				f.state = InMarker
				markerContent.Reset()
			} else if char == '\'' {
				// Handle quotes
				if currentWord.Len() > 0 {
					result.WriteString(currentWord.String())
					currentWord.Reset()
				}
				f.state = InQuotes
				result.WriteRune(char)
			} else if char == ' ' || char == '\t' || char == '\n' {
				// Word boundary
				if currentWord.Len() > 0 {
					result.WriteString(currentWord.String())
					currentWord.Reset()
				}
				result.WriteRune(char)
			} else {
				currentWord.WriteRune(char)
			}
			
		case InMarker:
			if char == ')' {
				// Process the marker and previous word
				marker := markerContent.String()
				f.state = Normal
				
				// Apply transformation based on marker
				prevText := result.String()
				transformedText := f.applyMarkerTransformation(prevText, marker)
				result.Reset()
				result.WriteString(transformedText)
			} else {
				markerContent.WriteRune(char)
			}
			
		case InQuotes:
			result.WriteRune(char)
			if char == '\'' {
				f.state = Normal
			}
		}
	}
	
	// Add any remaining word
	if currentWord.Len() > 0 {
		result.WriteString(currentWord.String())
	}
	
	// Clean up quotes and punctuation
	finalResult := result.String()
	finalResult = rules.CleanQuotes(finalResult)
	finalResult = rules.FixPunctuation(finalResult)
	
	return finalResult
}

// applyMarkerTransformation applies transformations based on markers
func (f *FSM) applyMarkerTransformation(text, marker string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}
	
	// Handle hex conversion
	if marker == "hex" {
		lastWord := words[len(words)-1]
		if val, err := strconv.ParseInt(lastWord, 16, 64); err == nil {
			words[len(words)-1] = strconv.FormatInt(val, 10)
		}
		return strings.Join(words, " ")
	}
	
	// Handle bin conversion
	if marker == "bin" {
		lastWord := words[len(words)-1]
		if val, err := strconv.ParseInt(lastWord, 2, 64); err == nil {
			words[len(words)-1] = strconv.FormatInt(val, 10)
		}
		return strings.Join(words, " ")
	}
	
	// Handle case transformations
	if marker == "up" {
		if len(words) > 0 {
			words[len(words)-1] = strings.ToUpper(words[len(words)-1])
		}
		return strings.Join(words, " ")
	}
	
	if marker == "low" {
		if len(words) > 0 {
			words[len(words)-1] = strings.ToLower(words[len(words)-1])
		}
		return strings.Join(words, " ")
	}
	
	if marker == "cap" {
		if len(words) > 0 {
			lastWord := words[len(words)-1]
			// Don't override already uppercase words
			if lastWord != strings.ToUpper(lastWord) || len(lastWord) == 1 {
				words[len(words)-1] = strings.Title(strings.ToLower(lastWord))
			}
		}
		return strings.Join(words, " ")
	}
	
	// Handle numbered transformations (simplified)
	if strings.Contains(marker, ",") {
		parts := strings.Split(marker, ",")
		if len(parts) == 2 {
			cmd := strings.TrimSpace(parts[0])
			nStr := strings.TrimSpace(parts[1])
			if n, err := strconv.Atoi(nStr); err == nil && n > 0 {
				if n <= len(words) {
					for i := len(words) - n; i < len(words); i++ {
						switch cmd {
						case "up":
							words[i] = strings.ToUpper(words[i])
						case "low":
							words[i] = strings.ToLower(words[i])
						case "cap":
							if words[i] != strings.ToUpper(words[i]) || len(words[i]) == 1 {
								words[i] = strings.Title(strings.ToLower(words[i]))
							}
						}
					}
				}
			}
		}
		return strings.Join(words, " ")
	}
	
	return text
}
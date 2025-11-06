package processor

import (
	"go-reloaded/internal/tokenizer"
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

// Process applies rules using FSM approach with tokenization
func (f *FSM) Process(text string) string {
	tokens := tokenizer.Tokenize(text)
	result := make([]string, 0, len(tokens))
	
	for i, token := range tokens {
		switch f.state {
		case Normal:
			if token.Type == tokenizer.Marker {
				f.state = InMarker
				result = f.applyMarker(result, token.Value)
			} else if strings.HasPrefix(token.Value, "'") {
				f.state = InQuotes
				result = append(result, f.cleanQuote(token.Value))
			} else {
				result = append(result, f.fixArticle(result, token.Value))
			}
		case InMarker:
			f.state = Normal
			result = append(result, token.Value)
		case InQuotes:
			if strings.HasSuffix(token.Value, "'") {
				f.state = Normal
			}
			result = append(result, token.Value)
		}
	}
	
	return f.fixPunctuation(strings.Join(result, " "))
}

func (f *FSM) applyMarker(result []string, marker string) []string {
	if len(result) == 0 {
		return result
	}
	
	lastIdx := len(result) - 1
	lastWord := result[lastIdx]
	
	if marker == "(hex)" {
		if val, err := strconv.ParseInt(lastWord, 16, 64); err == nil {
			result[lastIdx] = strconv.FormatInt(val, 10)
		}
	} else if marker == "(bin)" {
		if val, err := strconv.ParseInt(lastWord, 2, 64); err == nil {
			result[lastIdx] = strconv.FormatInt(val, 10)
		}
	} else if marker == "(up)" {
		result[lastIdx] = strings.ToUpper(lastWord)
	} else if marker == "(low)" {
		result[lastIdx] = strings.ToLower(lastWord)
	} else if marker == "(cap)" {
		result[lastIdx] = strings.Title(strings.ToLower(lastWord))
	}
	
	return result
}

func (f *FSM) fixArticle(result []string, word string) string {
	if len(result) > 0 && (result[len(result)-1] == "a" || result[len(result)-1] == "A") {
		if len(word) > 0 && strings.ContainsAny(string(word[0]), "aeiouAEIOUhH") {
			if result[len(result)-1] == "A" {
				result[len(result)-1] = "An"
			} else {
				result[len(result)-1] = "an"
			}
		}
	}
	return word
}

func (f *FSM) cleanQuote(text string) string {
	return strings.ReplaceAll(strings.ReplaceAll(text, "' ", "'"), " '", "'")
}

func (f *FSM) fixPunctuation(text string) string {
	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")
	return text
}
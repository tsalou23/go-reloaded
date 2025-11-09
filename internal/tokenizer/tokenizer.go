package tokenizer

import (
	"strings"
)

// TokenType represents different types of tokens
type TokenType int

const (
	Word TokenType = iota
	Punctuation
	Quote
	Marker
	Whitespace
)

// Token represents a parsed token
type Token struct {
	Type  TokenType
	Value string
}

// Tokenizer implements FSM-based tokenization
type Tokenizer struct {
	state int
}

// NewTokenizer creates a new tokenizer
func NewTokenizer() *Tokenizer {
	return &Tokenizer{state: 0}
}

// Tokenize parses text into tokens using FSM
func (t *Tokenizer) Tokenize(text string) []Token {
	var tokens []Token
	var current strings.Builder
	
	for _, char := range text {
		switch {
		case char == ' ' || char == '\t' || char == '\n':
			if current.Len() > 0 {
				tokens = append(tokens, Token{Type: Word, Value: current.String()})
				current.Reset()
			}
			tokens = append(tokens, Token{Type: Whitespace, Value: string(char)})
			
		case char == '\'' || char == '"':
			if current.Len() > 0 {
				tokens = append(tokens, Token{Type: Word, Value: current.String()})
				current.Reset()
			}
			tokens = append(tokens, Token{Type: Quote, Value: string(char)})
			
		case char == '(' || char == ')':
			if current.Len() > 0 {
				tokens = append(tokens, Token{Type: Word, Value: current.String()})
				current.Reset()
			}
			tokens = append(tokens, Token{Type: Marker, Value: string(char)})
			
		case char == ',' || char == '.' || char == '!' || char == '?' || char == ':' || char == ';':
			if current.Len() > 0 {
				tokens = append(tokens, Token{Type: Word, Value: current.String()})
				current.Reset()
			}
			tokens = append(tokens, Token{Type: Punctuation, Value: string(char)})
			
		default:
			current.WriteRune(char)
		}
	}
	
	if current.Len() > 0 {
		tokens = append(tokens, Token{Type: Word, Value: current.String()})
	}
	
	return tokens
}

// Reconstruct rebuilds text from tokens with proper spacing
func (t *Tokenizer) Reconstruct(tokens []Token) string {
	var result strings.Builder
	
	for i, token := range tokens {
		switch token.Type {
		case Word:
			result.WriteString(token.Value)
			
		case Punctuation:
			result.WriteString(token.Value)
			
		case Quote:
			result.WriteString(token.Value)
			
		case Marker:
			result.WriteString(token.Value)
			
		case Whitespace:
			// Only add whitespace if it's not redundant
			if i > 0 && i < len(tokens)-1 {
				prev := tokens[i-1]
				next := tokens[i+1]
				if prev.Type != Punctuation && next.Type != Punctuation {
					result.WriteString(token.Value)
				}
			}
		}
	}
	
	return result.String()
}

// PreprocessTokens applies smart preprocessing based on token analysis
func (t *Tokenizer) PreprocessTokens(tokens []Token) string {
	var result strings.Builder
	inQuote := false
	
	for i, token := range tokens {
		switch token.Type {
		case Word:
			result.WriteString(token.Value)
			
		case Punctuation:
			// Smart punctuation spacing
			if i > 0 && tokens[i-1].Type == Word {
				// No space before punctuation
				result.WriteString(token.Value)
			} else {
				result.WriteString(token.Value)
			}
			
		case Quote:
			// Smart quote handling
			if !inQuote {
				// Opening quote - remove space after
				result.WriteString(token.Value)
				inQuote = true
			} else {
				// Closing quote - remove space before
				result.WriteString(token.Value)
				inQuote = false
			}
			
		case Marker:
			result.WriteString(token.Value)
			
		case Whitespace:
			// Smart whitespace handling
			if i > 0 && i < len(tokens)-1 {
				prev := tokens[i-1]
				next := tokens[i+1]
				
				// Add space between words, but handle special cases
				if prev.Type == Word && next.Type == Word {
					result.WriteString(" ")
				} else if prev.Type == Word && next.Type == Marker {
					result.WriteString(" ")
				} else if prev.Type == Marker && next.Type == Word {
					result.WriteString(" ")
				}
			}
		}
	}
	
	return result.String()
}
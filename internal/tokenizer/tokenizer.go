package tokenizer

import "strings"

// Token represents a text token
type Token struct {
	Value string
	Type  TokenType
}

// TokenType defines token categories
type TokenType int

const (
	Word TokenType = iota
	Marker
	Punctuation
	Whitespace
)

// Tokenize splits text into tokens
func Tokenize(text string) []Token {
	var tokens []Token
	words := strings.Fields(text)
	
	for _, word := range words {
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			tokens = append(tokens, Token{Value: word, Type: Marker})
		} else {
			tokens = append(tokens, Token{Value: word, Type: Word})
		}
	}
	
	return tokens
}
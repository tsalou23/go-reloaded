package processor

import (
	"strings"
	"strconv"
)

// RealtimeFSM processes characters as they're typed
type RealtimeFSM struct {
	state       FSMState
	buffer      strings.Builder
	markerBuf   strings.Builder
	output      strings.Builder
	lastWord    string
}

// NewRealtimeFSM creates a new real-time FSM
func NewRealtimeFSM() *RealtimeFSM {
	return &RealtimeFSM{state: Normal}
}

// ProcessChar handles a single character input and returns any output
func (r *RealtimeFSM) ProcessChar(char rune) string {
	switch r.state {
	case Normal:
		return r.handleNormal(char)
	case InMarker:
		return r.handleMarker(char)
	case InQuotes:
		return r.handleQuotes(char)
	}
	return ""
}

func (r *RealtimeFSM) handleNormal(char rune) string {
	switch char {
	case '(':
		r.lastWord = r.buffer.String()
		r.buffer.Reset()
		r.state = InMarker
		r.markerBuf.Reset()
		return ""
		
	case ' ', '\t', '\n':
		word := r.buffer.String()
		r.buffer.Reset()
		if word != "" {
			r.output.WriteString(word)
		}
		r.output.WriteRune(char)
		result := r.output.String()
		r.output.Reset()
		return result
		
	case '\'':
		word := r.buffer.String()
		r.buffer.Reset()
		if word != "" {
			r.output.WriteString(word)
		}
		r.output.WriteRune(char)
		r.state = InQuotes
		return ""
		
	default:
		r.buffer.WriteRune(char)
		return ""
	}
}

func (r *RealtimeFSM) handleMarker(char rune) string {
	if char == ')' {
		marker := r.markerBuf.String()
		r.state = Normal
		
		// Apply transformation
		transformed := r.applyTransformation(r.lastWord, marker)
		r.output.WriteString(transformed)
		result := r.output.String()
		r.output.Reset()
		return result
	}
	
	r.markerBuf.WriteRune(char)
	return ""
}

func (r *RealtimeFSM) handleQuotes(char rune) string {
	r.output.WriteRune(char)
	if char == '\'' {
		r.state = Normal
		result := r.output.String()
		r.output.Reset()
		return result
	}
	return ""
}

func (r *RealtimeFSM) applyTransformation(word, marker string) string {
	switch marker {
	case "hex":
		if val, err := strconv.ParseInt(word, 16, 64); err == nil {
			return strconv.FormatInt(val, 10)
		}
	case "bin":
		if val, err := strconv.ParseInt(word, 2, 64); err == nil {
			return strconv.FormatInt(val, 10)
		}
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		return strings.Title(strings.ToLower(word))
	}
	return word
}

// GetCurrentBuffer returns current incomplete input
func (r *RealtimeFSM) GetCurrentBuffer() string {
	return r.buffer.String()
}

// Reset clears all state
func (r *RealtimeFSM) Reset() {
	r.state = Normal
	r.buffer.Reset()
	r.markerBuf.Reset()
	r.output.Reset()
	r.lastWord = ""
}
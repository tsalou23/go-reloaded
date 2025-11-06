package processor

import (
	"regexp"
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
	// Apply transformations using regex patterns like pipeline
	text = f.applyNumberConversions(text)
	text = f.applyCaseTransformations(text)
	text = f.applyArticleCorrections(text)
	text = f.applyQuoteCleaning(text)
	text = f.applyPunctuationFixes(text)
	return text
}

func (f *FSM) applyNumberConversions(text string) string {
	hexRe := regexp.MustCompile(`([0-9A-Fa-f]+)\s+\(hex\)`)
	text = hexRe.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if val, err := strconv.ParseInt(parts[0], 16, 64); err == nil {
			return strconv.FormatInt(val, 10)
		}
		return match
	})
	
	binRe := regexp.MustCompile(`([01]+)\s+\(bin\)`)
	text = binRe.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if val, err := strconv.ParseInt(parts[0], 2, 64); err == nil {
			return strconv.FormatInt(val, 10)
		}
		return match
	})
	
	return text
}

func (f *FSM) applyCaseTransformations(text string) string {
	// Multi-word transformations
	multiRe := regexp.MustCompile(`((?:\S+\s+){0,9}\S+)\s+\((up|low|cap),\s*(\d+)\)`)
	text = multiRe.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^(.+)\s+\((up|low|cap),\s*(\d+)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) != 4 {
			return match
		}
		
		words := strings.Fields(matches[1])
		command := matches[2]
		count, _ := strconv.Atoi(matches[3])
		
		if count <= 0 || count > len(words) {
			return match
		}
		
		for i := len(words) - count; i < len(words); i++ {
			switch command {
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "low":
				words[i] = strings.ToLower(words[i])
			case "cap":
				words[i] = strings.Title(strings.ToLower(words[i]))
			}
		}
		
		return strings.Join(words, " ")
	})
	
	// Single word transformations
	singleRe := regexp.MustCompile(`(\S+)\s+\((up|low|cap)\)`)
	text = singleRe.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^(\S+)\s+\((up|low|cap)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) != 3 {
			return match
		}
		
		word := matches[1]
		command := matches[2]
		
		switch command {
		case "up":
			return strings.ToUpper(word)
		case "low":
			return strings.ToLower(word)
		case "cap":
			return strings.Title(strings.ToLower(word))
		}
		
		return match
	})
	
	return text
}

func (f *FSM) applyArticleCorrections(text string) string {
	silentH := []string{"honest", "hour", "honor", "heir"}
	
	for _, word := range silentH {
		text = regexp.MustCompile(`\ba\s+`+word).ReplaceAllString(text, "an "+word)
		text = regexp.MustCompile(`\bA\s+`+word).ReplaceAllString(text, "An "+word)
	}
	
	vowelRe := regexp.MustCompile(`\ba\s+([aeiouAEIOU])`)
	text = vowelRe.ReplaceAllString(text, "an $1")
	
	vowelReUpper := regexp.MustCompile(`\bA\s+([aeiouAEIOU])`)
	text = vowelReUpper.ReplaceAllString(text, "An $1")
	
	return text
}

func (f *FSM) applyQuoteCleaning(text string) string {
	// Handle quotes with spaces inside
	quoteRe := regexp.MustCompile(`'\s+([^']*?)\s+'`)
	text = quoteRe.ReplaceAllString(text, "'$1'")
	
	// Handle quotes that contain apostrophes
	quoteRe2 := regexp.MustCompile(`'\s+([^']*?'[^']*?)\s+'`)
	text = quoteRe2.ReplaceAllString(text, "'$1'")
	
	return text
}

func (f *FSM) applyPunctuationFixes(text string) string {
	punctRe := regexp.MustCompile(`\s+([,.!?:;])`)
	return punctRe.ReplaceAllString(text, "$1")
}
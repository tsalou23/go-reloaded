package rules

import (
	"regexp"
)

// FixPunctuation fixes spacing around punctuation marks
func FixPunctuation(text string) string {
	// Remove spaces before punctuation marks and add space after comma
	punctRegex := regexp.MustCompile(`\s+([.,:;!?]+)`)
	text = punctRegex.ReplaceAllString(text, "$1")
	
	// Add space after comma if not followed by space
	commaRegex := regexp.MustCompile(`,([^\s])`)
	text = commaRegex.ReplaceAllString(text, ", $1")
	
	// Handle ellipsis and multiple punctuation
	ellipsisRegex := regexp.MustCompile(`\.\s*\.\s*\.`)
	text = ellipsisRegex.ReplaceAllString(text, "...")
	
	// Handle ?! combinations
	questExclRegex := regexp.MustCompile(`\?\s*!`)
	text = questExclRegex.ReplaceAllString(text, "?!")
	
	return text
}
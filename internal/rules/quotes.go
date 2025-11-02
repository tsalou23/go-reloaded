package rules

import (
	"regexp"
)

// CleanQuotes removes spaces inside single quotes
func CleanQuotes(text string) string {
	// Match content inside single quotes and remove internal spaces
	quoteRegex := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	
	return quoteRegex.ReplaceAllStringFunc(text, func(match string) string {
		// Extract content between quotes
		content := match[1 : len(match)-1] // Remove outer quotes
		// Trim spaces and return with quotes
		return "'" + regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(content, "") + "'"
	})
}
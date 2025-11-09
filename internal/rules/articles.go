package rules

import (
	"regexp"
	"strings"
)

// FixArticles changes "a" to "an" before vowels and silent h
func FixArticles(text string) string {
	// Handle silent h words first
	silentH := []string{"honest", "hour", "honor", "heir"}
	for _, word := range silentH {
		// Case insensitive matching for silent h words
		text = regexp.MustCompile(`\ba\s+`+word).ReplaceAllString(text, "an "+word)
		text = regexp.MustCompile(`\bA\s+`+word).ReplaceAllString(text, "An "+word)
		text = regexp.MustCompile(`\ba\s+`+strings.Title(word)).ReplaceAllString(text, "an "+strings.Title(word))
		text = regexp.MustCompile(`\bA\s+`+strings.Title(word)).ReplaceAllString(text, "An "+strings.Title(word))
	}
	
	// Handle vowel words with proper case preservation
	vowelRegex := regexp.MustCompile(`\b(a)\s+([aeiouAEIOU]\w*)`)
	text = vowelRegex.ReplaceAllString(text, "an $2")
	
	// For uppercase A, preserve the uppercase in "An"
	vowelRegexUpper := regexp.MustCompile(`\b(A)\s+([aeiouAEIOU]\w*)`)
	text = vowelRegexUpper.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			word := parts[1]
			// If next word is all uppercase, use "AN"
			if word == strings.ToUpper(word) && len(word) > 1 {
				return "AN " + word
			}
			// Otherwise use "An" to preserve the capital
			return "An " + word
		}
		return match
	})
	
	// Fix incorrect "an" before consonants (except silent h)
	consonantRegex := regexp.MustCompile(`\ban\s+([bcdfgjklmnpqrstvwxyzBCDFGJKLMNPQRSTVWXYZ]\w*)`)
	text = consonantRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			word := parts[1]
			// Check if it's a silent h word
			for _, silentWord := range silentH {
				if strings.EqualFold(word, silentWord) {
					return match // Keep "an" for silent h
				}
			}
			return "a " + word
		}
		return match
	})
	
	consonantRegexUpper := regexp.MustCompile(`\b(AN)\s+([bcdfgjklmnpqrstvwxyzBCDFGJKLMNPQRSTVWXYZ]\w*)`)
	text = consonantRegexUpper.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			word := parts[1]
			// Check if it's a silent h word
			for _, silentWord := range silentH {
				if strings.EqualFold(word, silentWord) {
					return match // Keep "AN" for silent h
				}
			}
			// Preserve uppercase: AN -> A
			return "A " + word
		}
		return match
	})
	
	return text
}
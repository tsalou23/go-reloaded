package rules

import (
	"regexp"
	"strings"
)

// FixArticles changes "a" to "an" before vowels and silent h
func FixArticles(text string) string {
	// Match "a" or "A" followed by word starting with vowel or h
	articleRegex := regexp.MustCompile(`\b[aA]\s+([aeiouAEIOUhH]\w*)`)
	
	return articleRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			if strings.HasPrefix(parts[0], "A") {
				return "An " + parts[1]
			}
			return "an " + parts[1]
		}
		return match
	})
}
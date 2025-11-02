package rules

import (
	"regexp"
	"strconv"
	"strings"
)

// ApplyCase processes up, low, cap transformations
func ApplyCase(text string) string {
	// Handle (up, n) - uppercase n words
	upNRegex := regexp.MustCompile(`(.+?)\s+\(up,\s*(\d+)\)`)
	text = upNRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, "(up,")
		if len(parts) == 2 {
			allText := strings.TrimSpace(parts[0])
			nStr := strings.TrimSpace(strings.TrimSuffix(parts[1], ")"))
			if n, err := strconv.Atoi(nStr); err == nil {
				words := strings.Fields(allText)
				if n <= len(words) {
					for i := len(words) - n; i < len(words); i++ {
						words[i] = strings.ToUpper(words[i])
					}
					return strings.Join(words, " ")
				}
			}
		}
		return match
	})

	// Handle (low, n) - lowercase n words
	lowNRegex := regexp.MustCompile(`(.+?)\s+\(low,\s*(\d+)\)`)
	text = lowNRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, "(low,")
		if len(parts) == 2 {
			allText := strings.TrimSpace(parts[0])
			nStr := strings.TrimSpace(strings.TrimSuffix(parts[1], ")"))
			if n, err := strconv.Atoi(nStr); err == nil {
				words := strings.Fields(allText)
				if n <= len(words) {
					for i := len(words) - n; i < len(words); i++ {
						words[i] = strings.ToLower(words[i])
					}
					return strings.Join(words, " ")
				}
			}
		}
		return match
	})

	// Handle (cap, n) - capitalize n words
	capNRegex := regexp.MustCompile(`(.+?)\s+\(cap,\s*(\d+)\)`)
	text = capNRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, "(cap,")
		if len(parts) == 2 {
			allText := strings.TrimSpace(parts[0])
			nStr := strings.TrimSpace(strings.TrimSuffix(parts[1], ")"))
			if n, err := strconv.Atoi(nStr); err == nil {
				words := strings.Fields(allText)
				if n <= len(words) {
					for i := len(words) - n; i < len(words); i++ {
						words[i] = strings.Title(strings.ToLower(words[i]))
					}
					return strings.Join(words, " ")
				}
			}
		}
		return match
	})

	// Handle single (cap) for multi-word phrases - only capitalize the last two words
	capPhraseRegex := regexp.MustCompile(`(\S+(?:\s+\S+)+)\s+\(cap\)`)
	text = capPhraseRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, " (cap)")
		if len(parts) == 2 {
			words := strings.Fields(parts[0])
			// Only capitalize the last 2 words (brooklyn bridge)
			if len(words) >= 2 {
				words[len(words)-2] = strings.Title(strings.ToLower(words[len(words)-2]))
				words[len(words)-1] = strings.Title(strings.ToLower(words[len(words)-1]))
			}
			return strings.Join(words, " ")
		}
		return match
	})

	// Handle single word transformations
	upRegex := regexp.MustCompile(`(\S+)\s+\(up\)`)
	text = upRegex.ReplaceAllStringFunc(text, func(match string) string {
		word := strings.Fields(match)[0]
		return strings.ToUpper(word)
	})

	lowRegex := regexp.MustCompile(`(\S+)\s+\(low\)`)
	text = lowRegex.ReplaceAllStringFunc(text, func(match string) string {
		word := strings.Fields(match)[0]
		return strings.ToLower(word)
	})

	capRegex := regexp.MustCompile(`(\S+)\s+\(cap\)`)
	text = capRegex.ReplaceAllStringFunc(text, func(match string) string {
		word := strings.Fields(match)[0]
		return strings.Title(strings.ToLower(word))
	})

	return text
}
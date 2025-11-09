package rules

import (
	"regexp"
	"strconv"
	"strings"
)

// ApplyNumbers processes hex and bin conversions
func ApplyNumbers(text string) string {
	// Handle (hex) conversions
	hexRegex := regexp.MustCompile(`(\w+)\s+\(hex\)`)
	text = hexRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) >= 2 {
			hexStr := parts[0]
			if val, err := strconv.ParseInt(hexStr, 16, 64); err == nil {
				return strconv.FormatInt(val, 10)
			}
			// Invalid hex - remove the (hex) marker
			return hexStr
		}
		return match
	})

	// Handle (bin) conversions
	binRegex := regexp.MustCompile(`(\w+)\s+\(bin\)`)
	text = binRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) >= 2 {
			binStr := parts[0]
			if val, err := strconv.ParseInt(binStr, 2, 64); err == nil {
				return strconv.FormatInt(val, 10)
			}
			// Invalid bin - remove the (bin) marker
			return binStr
		}
		return match
	})

	return text
}
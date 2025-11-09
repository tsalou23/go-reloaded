package rules

import (
	"regexp"
	"strconv"
	"strings"
)

// ApplyCase processes up, low, cap transformations
func ApplyCase(text string) string {
	// Handle multiple consecutive commands first
	text = handleMultipleCommands(text)
	// Handle sequential numbered commands
	text = handleSequentialCommands(text)
	
	// Handle (up, n) - uppercase n words before the command
	upNRegex := regexp.MustCompile(`((?:\S+\s+)*\S*)\s+\(up,\s*(\d+)\)`)
	text = upNRegex.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^(.*)\s+\(up,\s*(\d+)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) == 3 {
			allText := strings.TrimSpace(matches[1])
			nStr := strings.TrimSpace(matches[2])
			if n, err := strconv.Atoi(nStr); err == nil && n > 0 {
				words := strings.Fields(allText)
				if len(words) > 0 {
					// Transform min(n, available_words) words from the end
					startIdx := len(words) - n
					if startIdx < 0 {
						startIdx = 0
					}
					for i := startIdx; i < len(words); i++ {
						words[i] = strings.ToUpper(words[i])
					}
					return strings.Join(words, " ")
				}
			}
		}
		return match
	})

	// Handle (low, n) - lowercase n words before the command
	lowNRegex := regexp.MustCompile(`((?:\S+\s+)*\S*)\s+\(low,\s*(\d+)\)`)
	text = lowNRegex.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^(.*)\s+\(low,\s*(\d+)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) == 3 {
			allText := strings.TrimSpace(matches[1])
			nStr := strings.TrimSpace(matches[2])
			if n, err := strconv.Atoi(nStr); err == nil && n > 0 {
				words := strings.Fields(allText)
				if len(words) > 0 {
					// Transform min(n, available_words) words from the end
					startIdx := len(words) - n
					if startIdx < 0 {
						startIdx = 0
					}
					for i := startIdx; i < len(words); i++ {
						words[i] = strings.ToLower(words[i])
					}
					return strings.Join(words, " ")
				}
			}
		}
		return match
	})

	// Handle (cap, n) - capitalize n words before the command
	capNRegex := regexp.MustCompile(`((?:\S+\s+)*\S*)\s+\(cap,\s*(\d+)\)`)
	text = capNRegex.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^(.*)\s+\(cap,\s*(\d+)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) == 3 {
			allText := strings.TrimSpace(matches[1])
			nStr := strings.TrimSpace(matches[2])
			if n, err := strconv.Atoi(nStr); err == nil && n > 0 {
				words := strings.Fields(allText)
				if len(words) > 0 {
					// Transform min(n, available_words) words from the end
					startIdx := len(words) - n
					if startIdx < 0 {
						startIdx = 0
					}
					for i := startIdx; i < len(words); i++ {
						// Don't override already uppercase words
						if words[i] != strings.ToUpper(words[i]) || len(words[i]) == 1 {
							words[i] = strings.Title(strings.ToLower(words[i]))
						}
					}
					return strings.Join(words, " ")
				}
			}
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
		// Don't override already uppercase words
		if word == strings.ToUpper(word) && len(word) > 1 {
			return word
		}
		return strings.Title(strings.ToLower(word))
	})

	return text
}

// handleMultipleCommands processes consecutive commands like (cap) (up)
func handleMultipleCommands(text string) string {
	// Handle patterns like "word (cmd1) (cmd2)" where both are case commands
	multiCmdRegex := regexp.MustCompile(`(\S+)\s+\((up|low|cap)\)\s+\((up|low|cap)\)`)
	text = multiCmdRegex.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^(\S+)\s+\((up|low|cap)\)\s+\((up|low|cap)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) == 4 {
			word := matches[1]
			firstCmd := matches[2]
			secondCmd := matches[3]
			
			// Apply first command
			switch firstCmd {
			case "up":
				word = strings.ToUpper(word)
			case "low":
				word = strings.ToLower(word)
			case "cap":
				word = strings.Title(strings.ToLower(word))
			}
			
			// Apply second command
			switch secondCmd {
			case "up":
				return strings.ToUpper(word)
			case "low":
				return strings.ToLower(word)
			case "cap":
				return strings.Title(strings.ToLower(word))
			}
			
			return word
		}
		return match
	})
	
	// Handle mixed commands like "word (case) (hex)" - apply case, keep other command
	mixedCmdRegex := regexp.MustCompile(`(\S+)\s+\((up|low|cap)\)\s+\(([^)]+)\)`)
	text = mixedCmdRegex.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^(\S+)\s+\((up|low|cap)\)\s+\(([^)]+)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) == 4 {
			word := matches[1]
			caseCmd := matches[2]
			otherCmd := matches[3]
			
			// Apply case command
			switch caseCmd {
			case "up":
				word = strings.ToUpper(word)
			case "low":
				word = strings.ToLower(word)
			case "cap":
				word = strings.Title(strings.ToLower(word))
			}
			
			// Return with the other command preserved
			return word + " (" + otherCmd + ")"
		}
		return match
	})
	
	return text
}

// handleSequentialCommands processes sequential numbered commands like (cap,2) (low,3)
func handleSequentialCommands(text string) string {
	// Handle patterns like "words (cmd1,n1) (cmd2,n2)"
	seqCmdRegex := regexp.MustCompile(`((?:\S+\s+)*\S*)\s+\((up|low|cap),\s*(\d+)\)\s+\((up|low|cap),\s*(\d+)\)`)
	text = seqCmdRegex.ReplaceAllStringFunc(text, func(match string) string {
		re := regexp.MustCompile(`^((?:\S+\s+)*\S*)\s+\((up|low|cap),\s*(\d+)\)\s+\((up|low|cap),\s*(\d+)\)$`)
		matches := re.FindStringSubmatch(match)
		if len(matches) == 6 {
			allText := strings.TrimSpace(matches[1])
			firstCmd := matches[2]
			firstN, _ := strconv.Atoi(matches[3])
			secondCmd := matches[4]
			secondN, _ := strconv.Atoi(matches[5])
			
			words := strings.Fields(allText)
			if len(words) == 0 {
				return match
			}
			
			// Apply first command
			if firstN > 0 && firstN <= len(words) {
				startIdx := len(words) - firstN
				if startIdx < 0 {
					startIdx = 0
				}
				for i := startIdx; i < len(words); i++ {
					switch firstCmd {
					case "up":
						words[i] = strings.ToUpper(words[i])
					case "low":
						words[i] = strings.ToLower(words[i])
					case "cap":
						if words[i] != strings.ToUpper(words[i]) || len(words[i]) == 1 {
							words[i] = strings.Title(strings.ToLower(words[i]))
						}
					}
				}
			}
			
			// Apply second command (may overlap with first)
			if secondN > 0 && secondN <= len(words) {
				startIdx := len(words) - secondN
				if startIdx < 0 {
					startIdx = 0
				}
				for i := startIdx; i < len(words); i++ {
					switch secondCmd {
					case "up":
						words[i] = strings.ToUpper(words[i])
					case "low":
						words[i] = strings.ToLower(words[i])
					case "cap":
						if words[i] != strings.ToUpper(words[i]) || len(words[i]) == 1 {
							words[i] = strings.Title(strings.ToLower(words[i]))
						}
					}
				}
			}
			
			return strings.Join(words, " ")
		}
		return match
	})
	return text
}
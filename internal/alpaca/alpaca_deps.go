package alpaca

import "strings"

// This is likely an overestimation of tokens - but that's fine.
func CountTokens(text string) int {
	if text == "" {
		return 0
	}

	count := 0
	wordStart := -1
	specialChars := "()<>/-=!@#$%^&*"

	// Single pass through the text
	for i, char := range text {
		switch {
		case char == '\n' || char == ' ' || char == '\t' || strings.ContainsRune(specialChars, char):
			count++
			if wordStart != -1 {
				// Count word tokens
				count += (i - wordStart) / 4
				wordStart = -1
			}
		default:
			if wordStart == -1 {
				wordStart = i
			}
		}
	}

	// Handle last word if text doesn't end with space
	if wordStart != -1 {
		count += (len(text) - wordStart + 3) / 4
	}

	return count
}

package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	wordStart := -1 // Indicates the start index of the current word

	for i := 0; i < len(s); i++ {
		// if current character is a separator
		if isSeparator(s[i]) {
			// if we've encountered a word before, add it to the result
			if wordStart != -1 {
				words = append(words, s[wordStart:i])
			       wordStart = -1 // Reset wordStart
			}
		} else if wordStart == -1 {
			// Start of a new word
			wordStart = i
		}
		// If it's the last character and it's not a separator, add the last word
		if i == len(s)-1 && wordStart != -1 {
			words = append(words, s[wordStart:i+1])
		}
	}

	return words
}

func isSeparator(c byte) bool {
	// Define separators as space, tab, and newline
	return c == ' ' || c == '\t' || c == '\n'
}

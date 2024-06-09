package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	wordStart := -1

	for i := 0; i < len(s); i++ {
		if isSeparator(s[i]) {
			if wordStart != -1 {
				words = append(words, s[wordStart:i])
				wordStart = -1
			}
		} else if wordStart == -1 {
			wordStart = i
		}
		if i == len(s)-1 && wordStart != -1 {
			words = append(words, s[wordStart:i+1])
		}
	}
	return words
}

func isSeparator(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

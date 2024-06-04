package piscine

func Capitalize(s string) string {
	chars := []rune(s)
	wordStart := true

	for i, char := range chars {
		if wordStart && isLetter(char) {
			chars[i] = toUpper(char)
			wordStart = false
		} else {
			chars[i] = toLower(char)
		}

		if !isLetter(char) {
			wordStart = true
		} else {
			// Check if the previous character was not a letter
            		if i > 0 && !isLetter(chars[i-1]) {
               		 wordStart = true
		 	}
		}
	}

	return string(chars)
}

// isLetter checks if a rune is a letter
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// toUpper converts a lowercase letter to uppercase
func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

// toLower converts an uppercase letter to lowercase
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

package piscine

func Capitalize(s string) string {
	chars := []rune(s)
	wordStart := true

	for i, char := range chars {
		if wordStart && isAlphanumeric(char) {
			chars[i] = toUpper(char)
			wordStart = false
		} else if isAlphanumeric(char) {
			chars[i] = toLower(char)
		}

		if !isAlphanumeric(char) {
			wordStart = true
		}
	}

	return string(chars)
}

// isLetter checks if a rune is a letter
func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
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

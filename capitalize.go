package piscine

import "unicode"

func Capitalize(s string) string {
	// Convert the string to a slice of runes
	chars := []rune(s)
    	wordStart := true
    	// Loop through each character
    	for i, char := range chars {
        	// Capitalize the first letter of each word
        	if wordStart && unicode.IsLetter(char) {
            		chars[i] = unicode.ToUpper(char)
           		wordStart = false
        	} else {
            		chars[i] = unicode.ToLower(char)
        	}
        	// Update wordStart flag
        	if unicode.IsSpace(char) || !unicode.IsLetter(char) {
            		wordStart = true
        	}
    	}
    // Convert the slice of runes back to a string and return
    return string(chars)
}

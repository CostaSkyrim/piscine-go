package piscine

func LastRune(s string) rune {
	sentence := []rune(s)[len(s)-1]
	return sentence
}

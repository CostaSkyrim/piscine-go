package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n >= 0 && n < len(runes) {
		return runes[n]
	}
	return 0
}

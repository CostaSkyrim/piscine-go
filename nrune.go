package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n >= 0 && n < (len(runes)+1) {
		return runes[n-1]
	}
	return 0
}

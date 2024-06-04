package piscine

func ToLower(s string) string {
	result := make([]rune, len(s))
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			result[i] = c - ('A' - 'a')
		} else {
			result[i] = c
		}
	}
	return string(result)
}

package piscine

func ToUpper(s string) string {
	result := make([]rune, len(s))
	for i, c := range s {
		if c >= 'a' && c <= 'z' {
			result[i] = c - ('a' - 'A')
		} else {
			result[i] = c
		}
	}
	return string(result)
}

package piscine

func JumpOver(str string) string {
	result := ""
	for i, s := range str {
		if i%3 == 2 {
			result += string(s)
		}
	}
	return result + "\n"
}

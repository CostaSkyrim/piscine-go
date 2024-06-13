package piscine

func JumpOver(str string) string {
	result := ""
	for i, s := range str {
		if i%3 == 2 && i != len(str)-1 {
			result += string(s)
		}
	}
	return result + "\n"
}

package piscine

func StringToIntSlice(str string) []int {
	var wn []int
	for _, r := range str {
		wn = append(wn, int(r))
	}
	return wn
}

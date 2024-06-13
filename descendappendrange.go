package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}

	length := max - min
	result := []int{}

	for i := 0; i < length; i++ {
		result = append(result, max-i)
	}

	return result
}

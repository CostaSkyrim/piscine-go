package piscine

func Compact(ptr *[]string) int {
	tempSlice := *ptr
	count := 0

	for _, s := range tempSlice {
		if s != "" {
			tempSlice[count] = s
			count++
		}
	}
	*ptr = tempSlice[:count]

	return count
}

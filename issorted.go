package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	// Determine the sorting order (ascending or descending)
	order := 0
	for i := 0; i < len(a)-1; i++ {
		comparison := f(a[i], a[i+1])
		if comparison != 0 {
			order = comparison
			break
		}
	}

	// Check if the array is sorted according to the determined order
	for i := 0; i < len(a)-1; i++ {
		if (order > 0 && f(a[i], a[i+1]) < 0) || (order < 0 && f(a[i], a[i+1]) > 0) {
			return false
		}
	}
	return true
}

func compare(a, b int) int {
	return a - b
}

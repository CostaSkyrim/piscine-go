package piscine

func ShoppingListSort(slice []string) []string {
	// Insertion sort algorithm
	for i := 1; i < len(slice); i++ {
		current := slice[i]
		j := i - 1
		// Move elements of slice[0..i-1], that are greater than current,
		// to one position ahead of their current position
		for j >= 0 && len(slice[j]) > len(current) {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = current
	}
	return slice
}

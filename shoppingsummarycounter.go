package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store the count of each item
	itemCount := make(map[string]int)

	// Variable to hold the current item
	var currentItem string

	// Iterate over each character in the string
	for i := 0; i < len(str); i++ {
		char := str[i]

		// If the character is a space, it means we've reached the end of an item
		if char == ' ' {
			// If currentItem is not empty, add it to the map
			if currentItem != "" {
				itemCount[currentItem]++
				currentItem = ""
			}
		} else {
			// Otherwise, keep building the current item
			currentItem += string(char)
		}
	}

	// Add the last item to the map if there's any
	if currentItem != "" {
		itemCount[currentItem]++
	}
	return itemCount
}

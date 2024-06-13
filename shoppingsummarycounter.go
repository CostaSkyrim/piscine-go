package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store the count of each item
	itemCount := make(map[string]int)

	// Split the string into words (items) based on spaces
	items := strings.Fields(str)

	// Iterate over the items and count each one
	for _, item := range items {
		itemCount[item]++
	}

	return itemCount
}

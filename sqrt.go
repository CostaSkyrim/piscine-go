package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0 // Square root is not defined for negative numbers
	}
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i // Return the square root if it's a whole number
		}
	}
	return 0 // Return 0 if the square root is not a whole number
}

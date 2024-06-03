package piscine

import "math"

func IterativeFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= nb; i++ {
		result *= i
		if result < 0 || result > math.MaxInt32 {
			return 0
		}
	}
	return result
}

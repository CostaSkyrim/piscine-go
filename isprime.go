package piscine

import "math"

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb % 2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(nb)))
	for i := 3; i <= sqrtN; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

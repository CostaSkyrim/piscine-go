package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	const maxInt = int(^uint(0) >> 1)
	result := 1
	for i := 2; i <= nb; i++ {
		result *= i
		if result > maxInt/i {
			return 0
		}
	}
	return result
}

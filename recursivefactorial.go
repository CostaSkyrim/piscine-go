package piscine

func RecursiveFactorial(nb int) int {
	const maxInt = int(^uint(0) >> 1)
	if nb < 0 {
		return 0 // Factorial is not defined for negative numbers
	}
	if nb == 0 || nb == 1 {
		return 1 // 0! and 1! are both 1
	}
	return factorialHelper(nb, 1, maxInt)
}

func factorialHelper(nb, accumulator, maxInt int) int {
	if nb == 1 {
		return accumulator
	}
	if accumulator > maxInt/nb {
		return 0 // Overflow would occur
	}
	return factorialHelper(nb-1, accumulator*nb, maxInt)
}

package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	for !IsPrime(nb) {
		nb += 2
	}
	return nb
}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false // 1 and any number less than 1 are not prime
	}
	if nb == 2 {
		return true // 2 is the only even prime number
	}
	if nb%2 == 0 {
		return false // Any other even number is not a prime number
	}
	sqrtN := intSqrt(nb) // Calculate the square root of n using integer method
	for i := 3; i <= sqrtN; i += 2 { // Check only odd numbers from 3 to sqrt(n)
		if nb%i == 0 {
			return false // If n is divisible by any number, it is not prime
		}
	}
	return true // If no divisors are found, n is prime
}

package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	sqrtN := intSqrt(nb)
	for i := 3; i <= sqrtN; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func intSqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}
	x, y := nb, (nb+1)/2
	for y < x {
		x = y
		y = (nb/y + y) / 2
	}
	return x
}

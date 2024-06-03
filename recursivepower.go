package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * powerRecursive(nb, power-1)
}

func powerRecursive(nb int, power int) int {
	if power == 0 {
		return 1
	}
	return nb * powerRecursive(nb, power-1)
}

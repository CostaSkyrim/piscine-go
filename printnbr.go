package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	printChar := func(c rune) {
		z01.PrintRune(c)
	}
	if n < 0 {
		printChar('-')
		n = -n
	}
	if n == 0 {
		printChar('0')
		return
	}
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{'0' + rune(digit)}, digits...)
		n /= 10
	}
	for _, digit := range digits {
		printChar(digit)
	}
}

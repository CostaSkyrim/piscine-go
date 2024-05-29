package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	printChar := func(c rune) {
		z01.PrintRune(c)
	}
	if n == 0 {
		printChar('0')
		return
	}
	if value == -1<<31 {
		printChar('-')
		printChar('9')
		printChar('2')
		printChar('2')
		printChar('3')
		printChar('3')
		printChar('7')
		printChar('2')
		printChar('0')
		printChar('3')
		printChar('6')
		printChar('8')
		printChar('5')
		printChar('4')
		printChar('7')
		printChar('7')
		printChar('5')
		printChar('8')
		printChar('0')
		printChar('8')
		return
	}
	if n < 0 {
		printChar('-')
		n = -n
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

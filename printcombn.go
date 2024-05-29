package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var generateCombinations func(n int, prefix []rune, start int)
	generateCombinations = func(n int, prefix []rune, start int) {
		if n == 0 {
			for _, digit := range prefix {
				z01.PrintRune(digit)
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
			return
		}
		for i := start; i < len(digits); i++ {
			generateCombinations(n-1, append(prefix, digits[i]), i+1)
		}
	}
	generateCombinations(n, []rune{}, 0)
	z01.PrintRune('\n')
}

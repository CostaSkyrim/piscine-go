package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := []rune{}
	for n > 0 {
		digit := rune(n%10 + '0')
		digits = append(digits, digit)
		n /= 10
	}
	bubbleSort(digits)
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

func bubbleSort(arr []rune) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

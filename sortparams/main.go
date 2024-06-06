package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Skip the first argument, which is the program name

	// Bubble sort to sort the arguments in ASCII order
	n := len(args)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if args[j] > args[j+1] {
				// Swap args[j] and args[j+1]
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Iterate over each argument and print it
	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

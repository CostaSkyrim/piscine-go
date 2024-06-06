package main

import (
	"os"

	"sort"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Skip the first argument, which is the program name

	// Sort the arguments in ASCII order
	sort.Strings(args)

	// Iterate over each argument and print it
	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args:= os.Args[1:]

	for _, arg := range args {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

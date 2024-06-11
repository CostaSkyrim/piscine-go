package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := os.Args[1:]
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have a odd number of arguments"

	if isEven(len(s)) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

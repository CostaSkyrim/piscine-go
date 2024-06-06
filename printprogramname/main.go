package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	programPath := os.Args[0]
	lastSlash := -1
	for i := 0; i < len(programPath); i++ {
		if programPath[i] == '/' {
			lastSlash = i
		}
	}

	// Extract the program name from the full path
	var programName string
	if lastSlash == -1 {
		programName = programPath
	} else {
		programName = programPath[lastSlash+1:]
	}

	// Print the program name using z01
	for _, r := range programName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

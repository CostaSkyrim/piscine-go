package main

import (
	"fmt"
	"os"
)

func printHelp() {
	helpMessage := `Usage: [OPTIONS] STRING

    --insert  (or -i) STRING    This flag inserts the given STRING into the argument string.
    --order   (or -o)           This flag sorts the argument string in ASCII order.
    --help    (or -h)           This flag prints the help message.`
	fmt.Println(helpMessage)
}

func insertString(mainStr, insertStr string) string {
	return insertStr + mainStr
}

func orderString(str string) string {
	strArr := []rune(str)

	for i := 0; i < len(strArr); i++ {
		for j := i + 1; j < len(strArr); j++ {
			if strArr[i] > strArr[j] {
				strArr[i], strArr[j] = strArr[j], strArr[i]
			}
		}
	}

	return string(strArr)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	var (
		insertStr string
		order     bool
		str       string
	)

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--insert", "-i":
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		case "--order", "-o":
			order = true
		default:
			str = args[i]
		}
	}

	if insertStr != "" {
		str = insertString(str, insertStr)
	}

	if order {
		str = orderString(str)
	}

	fmt.Println(str)
}

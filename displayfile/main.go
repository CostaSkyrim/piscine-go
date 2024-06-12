package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else if len(args) == 1 {
		file, err := os.Open("quest8.txt")
		if err != nil {
			fmt.Printf(err.Error())
		}
		arr := make([]byte, 24)
		file.Read(arr)
		fmt.Printf(string(arr))
		file.Close()
	} else {
		fmt.Println("File name missing")
	}
}

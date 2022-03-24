package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) > 0 {
		if len(arg) > 1 {
			fmt.Println("Too many arguments")
		} else if len(arg) == 1 {
			fmt.Println("Almost there!!")
		}
	} else {
		fmt.Println("File name missing")
	}
}

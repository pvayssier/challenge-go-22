package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := len(arguments) - 1; i >= 1; i-- {
		for j := 0; j < len(arguments[i]); j++ {
			z01.PrintRune(rune(arguments[i][j]))
		}
		z01.PrintRune('\n')
	}
}

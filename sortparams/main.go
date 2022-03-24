package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	long := (len(arguments))
	for i := 0; i < long; i++ {
		m := i
		for j := i + 1; j < long; j++ {
			if arguments[j] < arguments[m] {
				m = j
			}
		}
		arguments[m], arguments[i] = arguments[i], arguments[m]
	}
	for k := 0; k <= long-1; k++ {
		for i := 0; i <= len(arguments[k])-1; i++ {
			z01.PrintRune(rune(arguments[k][i]))
		}
		z01.PrintRune('\n')
	}
}

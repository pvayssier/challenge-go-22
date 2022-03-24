package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	runes := []byte(s)
	result := 0
	grandeur := 1
	booleen := false
	for i := len(runes) - 1; i >= 0; i-- {
		if []byte(s)[i] == 45 && booleen == false && i == 0 {
			booleen = true
			result *= -1
		} else if []byte(s)[i] == 43 && booleen == false && i == 0 {
			result *= 1
			booleen = true
		} else if []byte(s)[i] < 48 || []byte(s)[i] > 59 {
			return 0
		} else {
			result += int(runes[i]-48) * grandeur
			grandeur *= 10
		}
	}
	return result
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) > 0 {
		if arguments[0] == "--upper" {
			for i := 1; i < len(arguments); i++ {
				a := Atoi(arguments[i])
				if a <= 0 {
					z01.PrintRune(rune(32))
				} else if a > 26 {
					z01.PrintRune(rune(32))
				} else {
					z01.PrintRune(rune(a) + 64)
				}
			}
		} else {
			for i := 0; i < len(arguments); i++ {
				a := Atoi(arguments[i])
				if a <= 0 {
					z01.PrintRune(rune(32))
				} else if a > 26 {
					z01.PrintRune(rune(32))
				} else {
					z01.PrintRune(rune(a) + 96)
				}
			}
		}
		z01.PrintRune('\n')
	}
}

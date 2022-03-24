package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	var verif string
	if nb >= 0 {
		verif = "F"
	} else {
		verif = "T"
	}
	z01.PrintRune(rune(verif[0]))
	z01.PrintRune('\n')
}

package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	eppeler := []byte(s)
	for ind := range s {
		z01.PrintRune(rune(eppeler[ind]))
	}
}

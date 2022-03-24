package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	tempint := n
	runes := []rune{}
	counter := 0
	for i := 10; tempint > 0; {
		tempint = tempint / i
		counter++
	}
	if n == 0 {
		counter++
	}
	grandeur := RecursivePower(10, counter-1)
	for j := 0; j < counter; j++ {
		runes = append(runes, rune((n/grandeur)+48))
		n = n % grandeur
		grandeur = grandeur / 10
	}
	trirunes := SortIntegerTable2(runes)
	for k := counter - 1; k >= 0; k-- {
		z01.PrintRune(trirunes[k])
	}
}

package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var t []int
	v := 1
	if n < 0 {
		v = -1
		z01.PrintRune('-')
	}
	for n != 0 {
		reste := n % 10 * v  // example: reste = 1234 % 10 * v = 4 x 1 = 4
		t = append(t, reste) // append add a la fin du tableau "t" le reste
		n = (n / 10)         // n = n/10; (123,4 = 123. because 1234/10 = exactly 123.
		// when we reach the last number, for instance 1. 1%10 = 1 and then 1/10 = 0.
		// The loop stops because it only continues if n !=0. )
	}
	for i := len(t) - 1; i >= 0; i-- { // on décrémente l'index i pour aller du dernier jusqu'au 1er index (0)
		z01.PrintRune(rune(t[i] + 48)) // +48 car si on veut imprimer le int "1" il faut additionner le décimal "1" à 48 pour retrouve le int "1"
	}
}

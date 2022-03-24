package piscine

import "github.com/01-edu/z01"

func PrintCombN() {
	cpt := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > i {
				for k := 0; k < 10; k++ {
					if k > j {
						if cpt > 0 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						cpt++
						z01.PrintRune(rune(i) + 48)
						z01.PrintRune(rune(j) + 48)
						z01.PrintRune(rune(k) + 48)
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

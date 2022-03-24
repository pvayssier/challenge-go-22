package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 48; i < 58; i++ {
		for j := 48; j < 58 && i < 58; j++ {
			for k := 48; k < 58 && i < 58; k++ {
				for l := 48; l < 58 && i < 58; l++ {
					if ((int(rune(i))-48)*10 + (int(rune(j)) - 48)) < ((int(rune(k))-48)*10 + (int(rune(l)) - 48)) {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(32))
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if i == 57 && j == 56 && k == 57 && l == 57 {
							z01.PrintRune('\n')
							return
						} else {
							z01.PrintRune(',')
							z01.PrintRune(rune(32))
						}
					}
				}
			}
		}
	}
}

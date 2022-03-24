package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	voy := []rune{}
	indvoy := []int{}
	if len(arguments) < 1 {
		z01.PrintRune('\n')
	} else {
		for i := 0; i < len(arguments); i++ {
			for j := 0; j < len(arguments[i]); j++ {
				if arguments[i][j] == 'a' || arguments[i][j] == 'A' || arguments[i][j] == 'e' || arguments[i][j] == 'E' || arguments[i][j] == 'u' || arguments[i][j] == 'U' || arguments[i][j] == 'o' || arguments[i][j] == 'O' || arguments[i][j] == 'i' || arguments[i][j] == 'I' {
					voy = append(voy, rune(arguments[i][j]))
					indvoy = append(indvoy, j)
				}
			}
			indvoy = append(indvoy, -1)
		}
		ln := 0
		for k := 0; k < len(arguments); k++ {
			ln += len(arguments[k])
		}
		for i, j := 0, len(voy)-1; i < j; i, j = i+1, j-1 {
			voy[i], voy[j] = voy[j], voy[i]
		}
		fmt.Println(voy, indvoy)
		if len(voy) > 0 {
			j := 0
			ind := 0
			for i := 0; i < ln; i++ {
				if indvoy[ind] < 0 {
					j += 1
					ind += 1
					fmt.Println("j+1")
				}
				fmt.Println(indvoy[ind], voy[ind])
				if ind == len(indvoy)-2 {
					ind = 0
					fmt.Println("ind = 0")
				}
				if i == indvoy[ind] {
					z01.PrintRune(voy[ind])
					ind += 1
					fmt.Println("print voy")
				} else {
					z01.PrintRune(rune(arguments[j][i]))
					fmt.Println("print cons")
				}
			}
		} else {
			for i := 0; i < len(arguments); i++ {
				for j := 0; j < len(arguments[i]); j++ {
					z01.PrintRune(rune(arguments[i][j]))
				}
				z01.PrintRune(32)
			}
		}
	}
}

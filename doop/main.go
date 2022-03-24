package main

import (
	"os"
)

func main() {
	args := os.Args
	ui := 0
	robert := 0
	var gerard int
	if len(args) != 4 {
	} else {
		if args[1] == "9223372036854775809" || args[3] == "9223372036854775809" {
		} else {
			if Atoi(args[1]) != 9223372036854775807 && Atoi(args[3]) != 9223372036854775807 && (args[2] == "+" || args[2] == "-" || args[2] == "*" || args[2] == "/" || args[2] == "%") {
				ui = Atoi(args[1])
				robert = Atoi(args[3])
				if (ui >= 9223372036854775807 || robert >= 9223372036854775807) || (ui <= -9223372036854775808 || robert <= -9223372036854775808) {
				} else {
					if args[2] == "+" {
						gerard = ui + robert
					} else if args[2] == "-" {
						gerard = ui - robert
					} else if args[2] == "*" {
						gerard = ui * robert
					} else if args[2] == "/" {
						if ui == 0 || robert == 0 {
							os.Stdout.WriteString("No division by 0")
						} else {
							gerard = ui / robert
							if gerard == 0 {
								os.Stdout.WriteString("0")
							}
						}
					} else if args[2] == "%" {
						if ui != 0 && robert != 0 {
							gerard = ui % robert
						} else {
							os.Stdout.WriteString("No modulo by 0")
						}
					}
					os.Stdout.WriteString(PrintNbr(gerard))
				}
			}
		}
	}
}

func Atoi(s string) int {
	isNega := false
	sChangeable := []byte(s)
	result := 0
	grandeur := 1
	if len(s) == 0 {
		return 0
	}
	if int(sChangeable[0]) == 45 {
		isNega = true
		sChangeable[0] = '0'
	}
	if int(sChangeable[0]) == 43 {
		sChangeable[0] = '0'
	}

	for i := len(s) - 1; i >= 0; i-- {
		if int(sChangeable[i]) < 47 || int(sChangeable[i]) > 58 {
			return 9223372036854775807
		} else {
			result = result + (int(sChangeable[i])-48)*grandeur
			grandeur = grandeur * 10
		}
	}
	if isNega {
		result *= -1
	}
	return result
}

func PrintNbr(n int) string {
	var s []int
	var ui string
	if n < 0 {
		n *= -1
		ui = ui + "-"
		temp := n
		for i := 0; temp > 0; i++ {
			s = append(s, temp%10)
			temp /= 10
		}
		for x := len(s) - 1; x >= 0; x-- {
			ui = ui + string(rune(s[x])+48)
		}
		ui = ui + "\n"
		return ui
	} else {
		temp := n
		for i := 0; temp > 0; i++ {
			s = append(s, temp%10)
			temp /= 10
		}
		for x := len(s) - 1; x >= 0; x-- {
			ui = ui + string(rune(s[x])+48)
		}
		ui = ui + "\n"
		return ui
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	chstr := []string{}
	indstr := 0
	if len(arguments) == 0 {
		chstr = append(chstr, "help")
		arguments = append(arguments, " ")
	}
	for i := 0; i < len(arguments); i++ {
		if (len(arguments) == 1 && arguments[0] == "-h") || (len(arguments) == 1 && arguments[0] == "--help") {
			chstr = append(chstr, "help")
		}
		if len(arguments[i]) >= 9 {
			if arguments[i][:9] == "--insert=" {
				chstr = append(chstr, "insert1")
			}
		}
		if len(arguments[i]) >= 3 {
			if arguments[i][:3] == "-i=" {
				chstr = append(chstr, "insert2")
			}
		}
		if len(arguments[i]) == 7 {
			if arguments[i][:7] == "--order" {
				chstr = append(chstr, "order")
			}
		}
		if arguments[i] == "-o" {
			chstr = append(chstr, "order")
		}
	}
	if len(chstr) >= 2 {
		indstr = len(arguments) - 1
	} else if len(chstr) == 1 {
		indstr = 1
	}
	for z := 0; z < len(chstr); z++ {
		if chstr[z] == "insert1" {
			a := arguments[z][9:]
			arguments[indstr] += a
		} else if chstr[z] == "insert2" {
			a := arguments[z][3:]
			arguments[indstr] += a
		} else if chstr[z] == "help" {
			strhelp := "--insert\n" + "  -i\n" + "\t This flag inserts the string into the string passed as argument.\n" + "--order\n" + "  -o\n" + "\t This flag will behave like a boolean, if it is called it will order the argument."
			indstr = 0
			arguments[0] = strhelp
		} else if chstr[z] == "order" {
			arguments[indstr] = SortParams(arguments[indstr])
		}
	}
	fmt.Println(arguments[indstr])
}

func SortParams(arguments string) string {
	long := (len(arguments))
	chstr := []string{}
	byt := []byte(arguments)
	var str string

	for i := 0; i < long; i++ {
		chstr = append(chstr, string(byt[i]))
	}

	for i := 0; i < long; i++ {
		m := i
		for j := i + 1; j < long; j++ {
			if chstr[j] < chstr[m] {
				m = j
			}
		}
		chstr[m], chstr[i] = chstr[i], chstr[m]
	}
	for i := 0; i < long; i++ {
		str += chstr[i]
	}
	return str
}

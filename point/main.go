package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(32)
	z01.PrintRune('=')
	z01.PrintRune(32)
	diz := points.x / 10
	points.x -= (diz * 10)
	z01.PrintRune(rune(diz) + 48)
	z01.PrintRune(rune(points.x) + 48)
	z01.PrintRune(',')
	z01.PrintRune(32)
	z01.PrintRune('y')
	diz = points.y / 10
	points.y -= (diz * 10)
	z01.PrintRune(32)
	z01.PrintRune('=')
	z01.PrintRune(32)
	z01.PrintRune(rune(diz) + 48)
	z01.PrintRune(rune(points.y) + 48)
	z01.PrintRune('\n')
}

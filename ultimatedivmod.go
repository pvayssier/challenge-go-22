package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a % *b
	*a = *a / *b
	*b = c
}

package piscine

func StrLen(s string) int {
	runes := []rune(s)
	cpt := 0
	for i := range runes {
		cpt = i + 1
	}
	return cpt
}

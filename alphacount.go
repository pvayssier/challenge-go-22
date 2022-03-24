package piscine

func AlphaCount(s string) int {
	cpt := 0
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] <= 122 && runes[i] >= 97 {
			cpt++
		} else if runes[i] <= 90 && runes[i] >= 65 {
			cpt++
		}
	}
	return cpt
}

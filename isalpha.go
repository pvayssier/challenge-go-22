package piscine

func IsAlpha(s string) bool {
	cpt := 0
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 97 && runes[i] <= 122 {
			cpt++
		} else if runes[i] >= 65 && runes[i] <= 90 {
			cpt++
		} else if runes[i] >= 48 && runes[i] <= 57 {
			cpt++
		}
	}
	if cpt == len(s) {
		return true
	} else {
		return false
	}
}

package piscine

func IsUpper(s string) bool {
	cpt := 0
	runes := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if runes[i] >= 65 && runes[i] <= 90 {
			cpt++
		}
	}
	if cpt == len(s) {
		return true
	} else {
		return false
	}
}

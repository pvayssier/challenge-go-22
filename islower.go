package piscine

func IsLower(s string) bool {
	cpt := 0
	runes := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if runes[i] >= 97 && runes[i] <= 122 {
			cpt++
		}
	}
	if cpt == len(s) {
		return true
	} else {
		return false
	}
}

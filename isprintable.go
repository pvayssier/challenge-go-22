package piscine

func IsPrintable(s string) bool {
	cpt := 0
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		if runes[i] >= 32 && runes[i] <= 126 {
			cpt++
		}
	}
	if cpt == len(s) {
		return true
	} else {
		return false
	}
}

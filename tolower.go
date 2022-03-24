package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		temp := runes[i]
		if runes[i] >= 65 && runes[i] <= 90 {
			runes[i] = (rune(temp + 32))
		}
	}
	return string(runes)
}

package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := 0; i <= len(runes)-1; i++ {
		temp := runes[i]
		if runes[i] >= 97 && runes[i] <= 122 {
			runes[i] = (rune(temp - 32))
		}
	}
	return string(runes)
}

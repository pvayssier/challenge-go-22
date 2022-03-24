package piscine

func Capitalize2(s string) string {
	runes := []rune(s)
	runes2 := []rune(s)
	temp := IsAlpha(ToUpper(string(runes[0])))
	for i := 1; i <= len(runes)-1; i++ {
		if runes[i] >= 65 && runes[i] <= 90 || runes[i] == 32 {
			runes2[i] = runes[i]
		} else if runes[i] >= 48 && runes[i] <= 57 {
			runes2[i] = runes[i]
		} else if temp == false {
			runes2[i] = rune(runes[i] - 32)
		} else {
			runes2[i-1] = rune(runes[i-1])
		}
		temp = IsAlpha(string(runes[i]))
	}
	return string(runes2)
}

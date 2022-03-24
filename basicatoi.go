package piscine

func BasicAtoi(s string) int {
	runes := []byte(s)
	result := 0
	grandeur := 1
	for i := len(runes) - 1; i >= 0; i-- {
		result += int(runes[i]-48) * grandeur
		grandeur *= 10
	}
	return result
}

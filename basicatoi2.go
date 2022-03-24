package piscine

func BasicAtoi2(s string) int {
	runes := []byte(s)
	result := 0
	grandeur := 1
	for i := len(runes) - 1; i >= 0; i-- {
		if []byte(s)[i] < 48 || []byte(s)[i] > 57 {
			return 0
		} else {
			result += int(runes[i]-48) * grandeur
			grandeur *= 10
		}
	}
	return result
}

package piscine

func Atoi(s string) int {
	runes := []byte(s)
	result := 0
	grandeur := 1
	booleen := false
	for i := len(runes) - 1; i >= 0; i-- {
		if []byte(s)[i] == 45 && booleen == false && i == 0 {
			booleen = true
			result *= -1
		} else if []byte(s)[i] == 43 && booleen == false && i == 0 {
			result *= 1
			booleen = true
		} else if []byte(s)[i] < 48 || []byte(s)[i] > 59 {
			return 0
		} else {
			result += int(runes[i]-48) * grandeur
			grandeur *= 10
		}
	}
	return result
}

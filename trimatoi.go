package piscine

func TrimAtoi(s string) int {
	runes := []byte(s)
	result := 0
	grandeur := 1
	booleen := false
	symbole := 1
	for i := len(runes) - 1; i >= 0; i-- {
		if []byte(s)[i] == 45 && booleen == false {
			booleen = true
			symbole *= -1
		} else if []byte(s)[i] == 43 && booleen == false {
			booleen = true
		} else if []byte(s)[i] >= 48 && []byte(s)[i] <= 57 {
			result += int(runes[i]-48) * grandeur
			grandeur *= 10
		}
		if []byte(s)[i] >= 48 && []byte(s)[i] <= 57 && booleen == true {
			symbole = 1
		}
	}
	return result * symbole
}

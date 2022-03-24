package piscine

func Any(f func(string) bool, a []string) bool {
	runes := []bool{}
	for i := 0; i < len(a); i++ {
		runes = append(runes, f(a[i]))
	}
	for j := 0; j < len(runes); j++ {
		if runes[j] {
			return true
		}
	}
	return false
}

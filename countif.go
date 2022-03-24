package piscine

func CountIf(f func(string) bool, a []string) int {
	runes := []bool{}
	for i := 0; i < len(a); i++ {
		runes = append(runes, f(a[i]))
	}
	cpt := 0
	for j := 0; j < len(runes); j++ {
		if runes[j] {
			cpt += 1
		}
	}
	return cpt
}

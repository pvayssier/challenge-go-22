package piscine

func Map(f func(int) bool, a []int) []bool {
	runes := []bool{}
	for i := 0; i < len(a); i++ {
		runes = append(runes, f(a[i]))
	}
	return runes
}

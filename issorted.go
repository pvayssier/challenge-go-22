package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascendant, descendant := true, true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			ascendant = false
		}
	}
	for j := 0; j < len(a)-1; j++ {
		if f(a[j], a[j+1]) > 0 {
			descendant = false
		}
	}
	return ascendant || descendant
}

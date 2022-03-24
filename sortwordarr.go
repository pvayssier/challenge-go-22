package piscine

func SortWordArr(a []string) {
	result := []string{}
	long := (len(a))
	for i := 0; i < long; i++ {
		m := i
		for j := i + 1; j < long; j++ {
			if a[j] < a[m] {
				m = j
			}
		}
		a[m], a[i] = a[i], a[m]
	}
	for k := len(a) - 1; k >= 0; k-- {
		result = append(result, a[k])
	}
	a = result
}

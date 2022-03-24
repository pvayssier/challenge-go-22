package piscine

func SortIntegerTable(table []int) {
	result := []int{}
	long := (len(table))
	for i := 0; i < long; i++ {
		m := i
		for j := i + 1; j < long; j++ {
			if table[j] < table[m] {
				m = j
			}
		}
		table[m], table[i] = table[i], table[m]
	}
	for k := len(table) - 1; k >= 0; k-- {
		result = append(result, table[k])
	}
	table = result
}

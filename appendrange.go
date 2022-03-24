package piscine

func AppendRange(min, max int) []int {
	table := []int{}
	if min >= max {
		return []int(nil)
	} else {
		for i := min; i < max; i++ {
			table = append(table, i)
		}
		return table
	}
}

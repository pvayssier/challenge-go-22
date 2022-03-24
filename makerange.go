package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	} else {
		ln := max - min
		table := make([]int, ln)
		j := 0
		for i := min; i < max; i++ {
			table[j] = i
			j++
		}
		return table
	}
}

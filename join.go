package piscine

func Join(strs []string, sep string) string {
	str := strs[0] + sep
	for i := 1; i <= len(strs)-2; i++ {
		str += strs[i] + sep
	}
	str += strs[len(strs)-1]
	return str
}

package piscine

func Split2(s, sep string) []string {
	var tab []string
	j := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) < len(s)-1 {
			if s[i:(i+len(sep))] == sep {
				if len(s[j:i+len(sep)]) > 0 {
					tab = append(tab, s[j-1:i+len(sep)])
				}
			}
			j = i + len(sep) - 1
		} else if i == len(s)-1 {
			if len(s[j:i+1]) > 0 {
				tab = append(tab, s[j:i+1])
			}
			j = i
		}
	}
	return []string(tab)
}

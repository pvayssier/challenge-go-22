package piscine

func SplitWhiteSpaces(s string) []string {
	j := 0
	cpt := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ' ' && s[i-1] != ' ' && i <= len(s)-1 {
			cpt++
		}
	}
	chstr := make([]string, cpt+1)
	for k := 0; k < len(s); k++ {
		if s[k] == ' ' && s[k-1] != ' ' && k <= len(s)-1 {
			j += 1
		} else {
			if s[k] != ' ' {
				chstr[j] += string(s[k])
			}
		}
	}
	return chstr
}

package piscine

func BasicJoin(elems []string) string {
	strelems := string(elems[0])
	for i := 1; i <= len(elems)-1; i++ {
		strelems += string(elems[i])
	}
	return strelems
}

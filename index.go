package piscine

func Index(s string, toFind string) int {
	bytes := []byte(s)
	bytet := []byte(toFind)
	cpt := 0
	var ibis int
	if len(s) <= 0 || len(toFind) <= 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if bytes[i] == bytet[0] {
			ibis = i
			for j := 0; j < len(toFind) && j < len(s)-i; j++ {
				if bytes[i+j] == bytet[j] {
					cpt += 0
				} else {
					cpt += 1
				}
			}
			if cpt == 0 {
				return ibis
			} else {
				return -1
			}
		}
	}
	return -1
}

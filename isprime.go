package piscine

func IsPrime(nb int) bool {
	booleen := false
	cpt := 1
	if nb <= 1 {
		return booleen
	} else {
		for i := 1; i*i <= nb; i++ {
			if nb%i == 0 {
				cpt++
			}
		}
		if cpt < 3 {
			booleen = true
		}
	}
	return booleen
}

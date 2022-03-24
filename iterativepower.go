package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if nb == 0 || power < 0 {
		return 0
	} else {
		result := nb
		for i := 0; i < power-1; i++ {
			result = result * nb
		}
		return result
	}
}

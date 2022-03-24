package piscine

func IterativeFactorial(nb int) int {
	result := nb
	if nb < 0 || nb >= 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := nb - 1; i > 0; i-- {
			result = result * i
		}
		return result
	}
}

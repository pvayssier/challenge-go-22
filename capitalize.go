package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	cpt := 1
	var ui string
	for i := 0; i < len(s); i++ {
		if (runes[i] < 97 && runes[i] > 90) || (runes[i] < 65 && runes[i] > 57) || (runes[i] < 48 && runes[i] > 0) || runes[i] > 122 {
			cpt = cpt + 1
			ui = ui + string(runes[i])
		} else {
			if cpt > 0 {
				if runes[i] > 96 && runes[i] < 123 {
					runes[i] = runes[i] - 32
				}
				cpt = 0
				ui = ui + string(runes[i])
			} else {
				if runes[i] > 64 && runes[i] < 91 {
					runes[i] = runes[i] + 32
				}
				ui = ui + string(runes[i])
			}
		}
	}
	return ui
}

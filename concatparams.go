package piscine

func ConcatParams(args []string) string {
	var s string
	for i := 0; i < len(args); i++ {
		if i > 0 {
			s += "\n" + args[i]
		} else {
			s += args[i]
		}
	}
	return s
}

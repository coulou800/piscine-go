package piscine

func BasicAtoi(s string) int {
	n := 0
	f := 1

	for i := len(s) - 1; i >= 0; i-- {
		n += int(rune(s[i])-48) * f
		f *= 10
	}

	return n
}

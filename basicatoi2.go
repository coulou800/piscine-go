package piscine

func BasicAtoi2(s string) int {
	runes := []rune(s)
	n := 0
	for i := 0; i < len(s); i++ {
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		} else {
			n *= 10
			n += int(runes[i]) - '0'

		}
	}
	return n
}

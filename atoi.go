package piscine

func Atoi(s string) int {
	runes := []rune(s)
	n := 0
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			x := runes[i] != '+'
			if runes[i] < '0' || runes[i] > '9' {
				x := runes[0] != '-'
				if i != 0 && (x || runes[0] != '+') {
					return 0
				}
			} else if x || runes[i] != '-' {
				n *= 10
				n += int(runes[i]) - '0'
			}
		}
		if s[0] == '-' {
			return -n
		}
	}

	return n
}

package piscine

func IsPrintable(s string) bool {
	isprintable := false
	for i := range s {
		if rune(s[i]) >= rune(0) && rune(s[i]) <= rune(31) {
			isprintable = false
			break
		} else {
			isprintable = true
		}
	}
	return isprintable
}

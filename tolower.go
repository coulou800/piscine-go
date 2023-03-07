package piscine

func ToLower(s string) string {
	lower := []rune(s)
	for i := 0; i < len(lower); i++ {
		if IsUpper(string(s[i])) {
			lower[i] = (lower[i] - 'A') + 'a'
		}
	}
	return string(lower)
}

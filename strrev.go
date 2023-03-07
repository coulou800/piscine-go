package piscine

func StrRev(s string) string {
	runes := []rune(s)
	revRune := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		revRune = append(revRune, runes[i])
	}
	str := string(revRune)

	return str
}

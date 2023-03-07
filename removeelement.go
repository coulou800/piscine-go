package piscine

func RemoveElement(s []rune, i int) []rune {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

package piscine

func RemoveStringElem(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
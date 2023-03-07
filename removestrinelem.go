package piscine

func RemoveStringElem(s []string, i int) []string {
	s = append(s[:i],s[i+1:]...)
	return s[:1]
}
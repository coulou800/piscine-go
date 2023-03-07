package piscine

func RemoveStringElem(s []string, i int) []string {
	s = append(s[:i],s[i+1:]...)
	println(s)
	return s[:len(s)-1]
}
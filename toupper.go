package piscine

func ToUpper(s string) string {
	toupper := []rune(s)
	for i := 0; i <= StrLen(s)-1; i++ {
		if toupper[i] >= 'a' && toupper[i] <= 'z' {
			toupper[i] = toupper[i] - ' '
		}
	}
	return string(toupper)
}

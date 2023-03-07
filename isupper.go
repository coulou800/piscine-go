package piscine

func IsUpper(s string) bool {
	isupper := false
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			isupper = true
		} else {
			isupper = false
			break
		}
	}
	return isupper
}

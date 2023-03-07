package piscine

func IsAlpha(s string) bool {
	for i := range s {
		if s[i] < 'a' && s[i] > 'z' || s[i] < 'A' && s[i] > 'Z' || s[i] < '0' && s[i] > '9' {
			return false
		}
	}
	return true
}

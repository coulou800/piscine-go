package piscine

func IsLower(s string) bool {
	islower := false
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			islower = true
		} else {
			islower = false
			break
		}
	}
	return islower
}

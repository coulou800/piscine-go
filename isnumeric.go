package piscine

func IsNumeric(s string) bool {
	isnumeric := false
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			isnumeric = true
		} else {
			isnumeric = false
			break
		}
	}
	return isnumeric
}

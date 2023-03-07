package piscine

func Index(s string, toFind string) int {
	srunes := []rune(s)
	trunes := []rune(toFind)

	for i := 0; i < len(srunes)-len(trunes); i++ {
		if toFind == s[i:i+len(trunes)] {
			return i
		}
	}
	return -1
}

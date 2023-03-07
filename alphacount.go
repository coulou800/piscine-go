package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	count := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' || runes[i] >= 'A' && runes[i] <= 'Z' {
			count += 1
		}
	}
	return count
}

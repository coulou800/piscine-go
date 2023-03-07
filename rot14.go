package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, v := range runes {
		if v <= 'z' && v >= 'a' {
			if v <= 'l' {
				runes[i] += 14
			}
			if v > 'l' {
				runes[i] = 'a' + 14 - ('z' - runes[i]) - 1
			}
		} else if v <= 'Z' && v >= 'A' {
			if v <= 'L' {
				runes[i] += 14
			}
			if v > 'L' {
				runes[i] = 'A' + 14 - ('Z' - runes[i]) - 1
			}
		} else {
			continue
		}
	}
	return string(runes)
}

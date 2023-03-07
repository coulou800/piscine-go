package piscine

func SortWordArr(a []string) {
	sortAscii(a)
}

func sortAscii(b []string) {
	for i := 1; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
}

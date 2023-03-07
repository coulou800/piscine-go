package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(string(a[i]), string(a[j])) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

package piscine

func checker(a, b int) int {
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	asc := true
	desc := true
	for i, v := range a {
		if i < len(a)-1 {
			if checker(v, a[i+1]) > 0 {
				asc = false
			}
			if checker(v, a[i+1]) < 0 {
				desc = false
			}
		}
	}
	return asc || desc
}

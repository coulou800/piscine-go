package piscine

func Unmatch(a []int) int {
	for i, v := range a {
		found := false
		for j := 0; j < len(a); j++ {
			if v == a[j] && j != i {
				found = !found
			}
		}
		if !found {
			return v
		}
	}
	return -1
}

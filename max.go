package piscine

func Max(a []int) int {
	max := a[0]
	if len(a) > 0 {
		for i, v := range a {
			if i < len(a)-1 {
				if v > max {
					max = v
				}
			}
		}
		return max
	}
	return 0
}

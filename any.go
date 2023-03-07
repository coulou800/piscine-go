package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(string(v)) {
			return true
		}
	}
	return false
}

package piscine

func BasicJoin(elems []string) string {
	rst := ""
	for i := range elems {
		if i > 0 || i != len(elems)-1 {
			rst += elems[i]
		} else {
			rst += elems[i]
		}
	}
	return rst
}

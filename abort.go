package piscine

func Abort(a, b, c, d, e int) int {
	intarr := []int{a, b, c, d, e}
	SortIntegerTable(intarr)
	return intarr[len(intarr)/2]
}

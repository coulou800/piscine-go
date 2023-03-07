package piscine

func Compact(ptr *[]string) int {
	c := 0
	var b []string
	for _, v := range *ptr {
		if v != "" {
			b = append(b, v)
			c++
		}
	}
	*ptr = b
	return c
}

package piscine

func ActiveBits(n int) int {
	c := 0
	for i := 0; i < 32; i++ {
		if n>>i&1 == 1 {
			c++
		}
	}
	return c
}

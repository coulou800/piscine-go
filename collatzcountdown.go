package piscine

func CollatzCountdown(start int) int {
	var c int

	if start > 0 {
		for start > 1 {
			if start%2 == 0 {
				start /= 2
			} else {
				start *= 3
				start++
			}
			c++
		}
		return c
	}

	return -1
}

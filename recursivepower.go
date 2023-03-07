package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 || nb > 20 {
		return 0
	}
	if power == 0 {
		return 1
	}
	rst := nb * RecursivePower(nb, power-1)
	return rst
}

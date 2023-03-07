package piscine

func IterativeFactorial(n int) int {
	rst := 1
	if n < 0 {
		return 0
	} else if n == 0 || n == 1 {
		return 1
	} else if n > 20 {
		return 0
	} else {
		for i := 1; i <= n; i++ {
			rst *= i
		}
	}
	return rst
}

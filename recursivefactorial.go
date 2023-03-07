package piscine

func RecursiveFactorial(n int) int {
	if n < 0 || n > 20 {
		return 0
	}
	if n == 1 || n == 0 {
		return 1
	}
	return n * RecursiveFactorial(n-1)
}

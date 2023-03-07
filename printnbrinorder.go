package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	slice := []int{}
	if n == 0 {
		z01.PrintRune(rune(48))
	}
	for n > 0 {
		slice = append(slice, n%10)
		n = n / 10
	}
	SortIntegerTable(slice)

	for i := 0; i < len(slice); i++ {
		z01.PrintRune(rune(48 + slice[i]))
	}
}

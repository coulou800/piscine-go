package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	issignative := n < 0
	if issignative {
		n = -1 * n
		z01.PrintRune(rune(45))
	} else if n == 0 {
		z01.PrintRune(rune(48))
	}
	slice := []uint{}
	b := uint(n)
	for b > 0 {
		slice = append(slice, b%10)
		b = b / 10
	}
	for i, j := 0, len(slice)-1; i <= j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	for i := 0; i < len(slice); i++ {
		z01.PrintRune(rune(48 + slice[i]))
	}
}

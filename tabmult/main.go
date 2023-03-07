package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	num := Atoi(args[0])

	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune(i + 48))
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		for _, v := range args[0] {
			z01.PrintRune(v)
		}
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		printnbr(num * i)
		z01.PrintRune('\n')

	}
}

func Atoi(s string) int {
	num := 0
	r := []rune(s)

	for _, v := range r {
		num *= 10
		num += int(v - 48)
	}
	return num
}

func printnbr(n int) {
	var rst []rune
	for n > 0 {
		char := (n % 10) + 48
		rst = append(rst, rune(char))
		n /= 10
	}
	for i := len(rst) - 1; i >= 0; i-- {
		z01.PrintRune(rst[i])
	}
}

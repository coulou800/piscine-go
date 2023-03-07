package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		num, err := strconv.Atoi(os.Args[1])

		if err != nil {
			msg := "ERROR"
			for _, v := range msg {
				z01.PrintRune(v)
			}
		} else {
			hexConverter(num)
			z01.PrintRune('\n')
		}
	}
}

func hexConverter(n int) {
	base := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	res := []rune{}
	if n != 0 {
		res = append(res, base[n%16])
		n /= 16
		hexConverter(n)
	}
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(res[i])
	}
}

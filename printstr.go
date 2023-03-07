package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	str := []rune(s)
	i := 0
	for i < len(str) {
		z01.PrintRune(str[i])
		i++
	}
}

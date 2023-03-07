package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		runes := []rune(arg)
		for i, v := range runes {
			if isUpper(v) {
				runes[i] = v + 32
			} else if v == ' ' && i > 0 {
				runes[i-1] = runes[i-1] - 32
			}
			if !isUpper(runes[len(runes)-1]) && runes[len(runes)-1] != ' '{
				runes[len(runes)-1] = runes[len(runes)-1] - 32
			}
		}
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')

	}
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

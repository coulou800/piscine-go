package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	caps := false
	if len(args) != 0 {
		for _, v := range args {
			if v == "--upper" {
				caps = true
			}
		}
		if len(args) == 0 {
			z01.PrintRune(rune(0))
			return
		}
		for _, arg := range args {
			numv := 0
			for _, v := range arg {
				numv = numv*10 + int(v-'0')
			}
			if numv >= 1 && numv <= 26 {
				if !caps {
					z01.PrintRune(rune(numv + 96))
				} else {
					z01.PrintRune(rune(numv + 64))
				}
			} else if numv < '0' && numv > '9' {
				z01.PrintRune(' ')
			} else if numv > 24 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

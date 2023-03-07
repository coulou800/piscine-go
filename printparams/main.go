package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for j, i := range os.Args {
		if j != 0 {
			for _, r := range i {
				z01.PrintRune(r)
			}
		}
		if j > 0 && j < len(os.Args)-1 {
			z01.PrintRune('\n')
		}

	}
	z01.PrintRune('\n')
}

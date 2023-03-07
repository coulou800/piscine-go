package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, i := range os.Args[0] {
		if i != '/' && i != '.' {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}

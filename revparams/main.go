package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}

		if i > 1 && i <= len(os.Args)-1 {
			z01.PrintRune('\n')
		}

	}
	z01.PrintRune('\n')
}

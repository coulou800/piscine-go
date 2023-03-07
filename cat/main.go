package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		reader := os.Stdin
		txt, _ := io.ReadAll(reader)
		PrintStr(string(txt))

	} else {
		for _, v := range args {
			file, error := os.ReadFile(v)
			errorMsg := "ERROR: open " + v + ": no such file or directory\n"
			if error != nil {
				for _, g := range errorMsg {
					z01.PrintRune(rune(g))
				}
				os.Exit(1)
			} else {
				PrintStr(string(file))
			}
		}
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var c int
	var error bool

	for i, v := range args {
		if v == "-c" {
			c = Atoi(args[i+1])
		}
		if i >= 2 {
			file, err := os.ReadFile(v)
			if err != nil {
				error = true
				fmt.Printf("open %v: no such file or directory\n", v)
			} else {
				var text string
				if len(file) < c {
					text = string(file)
				} else {
					text = string(file)[len(file)-c:]
				}
				if len(args[2:]) == 1 {
					fmt.Printf("%v", text)
				} else {
					if i >= 3 {
						fmt.Printf("\n")
					}
					fmt.Printf("==> %v <==\n%v", v, text)
				}
			}
		}
	}
	if error {
		os.Exit(1)
	}
}

func Atoi(s string) int {
	runes := []rune(s)
	n := 0
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			x := runes[i] != '+'
			if runes[i] < '0' || runes[i] > '9' {
				x := runes[0] != '-'
				if i != 0 && (x || runes[0] != '+') {
					return 0
				}
			} else if x || runes[i] != '-' {
				n *= 10
				n += int(runes[i]) - '0'
			}
		}
		if s[0] == '-' {
			return -n
		}
	}

	return n
}

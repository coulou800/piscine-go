package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	options := []string{"--insert", "-i", "--order", "-o", "--help", "-h"}
	mustOrder := false
	strings := ""

	if len(args) > 0 {
		if args[0] != options[4] && args[0] != options[5] {
			strings = insert(strings, args[len(args)-1])
		}
	} else {
		printHelp()
	}

	for _, arg := range args {
		if len(arg) > len(options[0]) {
			if arg[:8] == options[0] {
				strings = insert(strings, arg[9:])
			}
		}
		if len(arg) > len(options[1]) {
			if arg[:2] == options[1] {
				strings = insert(strings, arg[3:])
			}
		}
		if arg == options[2] || arg == options[3] {
			mustOrder = true
		}
		if arg == options[4] || arg == options[5] {
			printHelp()
		}
	}
	if mustOrder {
		strings = sort(strings)
	}
	fmt.Println(strings)
}

func sort(s string) string {
	runes := []rune(s)
	for i := range string(s) {
		for j := range s {
			if runes[i] < runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func insert(s1, s2 string) string {
	r1 := []rune(s1)
	for _, s := range s2 {
		r1 = append(r1, s)
	}
	return string(r1)
}

func printHelp() {
	msg := `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`
	fmt.Print(msg)
}

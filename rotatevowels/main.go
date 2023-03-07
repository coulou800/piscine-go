package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	length := 0
	for i := range arguments {
		length = i + 1
	}
	if length != 0 {
		str := ""
		first := false
		for _, arg := range arguments {
			if first {
				str += " "
			}
			first = true
			str += arg
		}
		runes := []rune(str)
		var pos []int
		var v []rune
		for i, r := range runes {
			if isVowel(r) {
				pos = append(pos, i)
				v = append(v, runes[i])
			}
		}
		swap(v)
		for i := range pos {
			runes[pos[i]] = v[i]
		}
		for _, r := range runes {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

func swap(v []rune) {
	len := 0
	for i := range v {
		len = i + 1
	}
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		t := v[i]
		v[i] = v[j]
		v[j] = t
	}
}

func isVowel(a rune) bool {
	return (a == 'a' || a == 'e' || a == 'u' || a == 'o' || a == 'i' || a == 'A' || a == 'E' || a == 'U' || a == 'O' || a == 'I')
}

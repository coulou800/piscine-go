package hardonexam

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 || base[0] == base[1] || isIn(base, '-') || isIn(base, '+') {
		for _, v := range "NV" {
			z01.PrintRune(v)
		}
	} else {
		isNeg := nbr < 0

		if isNeg {
			nbr *= -1
			z01.PrintRune('-')
		}
		baseConverter(nbr, base)
	}

	z01.PrintRune('\n')
}

func isIn(base string, r rune) bool {
	for _, v := range base {
		if v == r {
			return true
		}
	}
	return false
}

func baseConverter(n int, base string) {
	lenbase := len(base)
	rst := []rune{}
	if n != 0 {
		rst = append(rst, []rune(base)[n%lenbase])
		n /= lenbase
		baseConverter(n, base)
	}
	for _, v := range rst {
		z01.PrintRune(v)
	}
}

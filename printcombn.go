package piscine

import "github.com/01-edu/z01"

func afficher(n int, digits [9]int, dmx [9]int) {
	i := 0
	for i < n {
		z01.PrintRune(rune(48 + digits[i]))
		i++
	}
	if digits[0] != dmx[0] {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func print_n1() {
	digits := [9]int{0}
	dmx := [9]int{9}
	for digits[0] <= dmx[0] {
		afficher(1, digits, dmx)
		digits[0]++
	}
}

func PrintCombN(n int) {
	digits := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	dmx := [9]int{}

	if n == 1 {
		print_n1()
	} else {
		i := n - 1
		j := 9
		for i >= 0 {
			dmx[i] = j
			i--
			j--
		}
		i = n - 1
		for digits[0] < dmx[0] {
			if digits[i] != dmx[i] {
				afficher(n, digits, dmx)
				digits[i]++
			}
			if digits[i] == dmx[i] {
				if digits[i-1] != dmx[i-1] {
					afficher(n, digits, dmx)
					digits[i-1]++
					j = i
					for j < n {
						digits[j] = digits[j-1] + 1
						j++
					}
					i = n - 1
				}
			}
			for digits[i] == dmx[i] && digits[i-1] == dmx[i-1] && i > 1 {
				i--
			}
		}
		afficher(n, digits, dmx)
	}
	z01.PrintRune('\n')
}

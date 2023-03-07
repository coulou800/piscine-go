package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i != j && j != k && i < j && j < k {
					if i != 0 || j != 1 || k != 2 {
						z01.PrintRune(' ')
					}
					z01.PrintRune(rune(48 + i))
					z01.PrintRune(rune(48 + j))
					z01.PrintRune(rune(48 + k))
					if i != 7 || j != 8 || k != 9 {
						z01.PrintRune(',')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

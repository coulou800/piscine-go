package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	grille := make([][]rune, len(args))
	for i := range args {
		grille[i] = []rune(args[i])
	}
	if isCorrect(grille) {
		for i := range grille {
			for k, j := range grille[i] {
				z01.PrintRune(j)
				if k < len(grille[i])-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func isValid(x, y, n int, grille [][]rune) bool {
	for i := range grille {
		if int(grille[y][i])-48 == n {
			return false
		}
	}
	for i := range grille {
		if int(grille[i][x])-48 == n {
			return false
		}
	}
	xC := (x / 3) * 3
	yC := (y / 3) * 3
	for i := yC; i < yC+3; i++ {
		for j := xC; j < xC+3; j++ {
			if int(grille[i][j])-48 == n {
				return false
			}
		}
	}
	return true
}

func isCorrect(g [][]rune) bool {
	filled := true
	for i := range g {
		for j := range g[i] {
			if g[i][j] == '.' {
				filled = false
			}
		}
	}
	if filled {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if g[y][x] == '.' {
				for z := 1; z <= 9; z++ {
					if isValid(x, y, z, g) {
						g[y][x] = rune(z + 48)
						if isCorrect(g) {
							return true
						}
					}
					g[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}

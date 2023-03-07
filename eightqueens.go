package piscine

import (
	"github.com/01-edu/z01"
)

func isSafePosition(lastPosition, testedPosition int, queenPosition [8]int) bool {
	for i := 0; i < lastPosition; i++ {
		position := queenPosition[i]
		// if the testedPosition is already taken return false
		// you attempt to place the queen on the same column
		if position == testedPosition {
			return false
		}
		// if the testedPosition is already taken return false
		if position == testedPosition-lastPosition+i {
			return false
		}
		if position == testedPosition+lastPosition-i {
			return false
		}
	}
	return true
}

func solve(indexPosition int, queenPosition [8]int) {
	if indexPosition == 8 {
		printSolution(queenPosition)
	} else {
		for i := 0; i < 8; i++ {
			if isSafePosition(indexPosition, i, queenPosition) {
				queenPosition[indexPosition] = i
				solve(indexPosition+1, queenPosition)
			}
		}
	}
}

func printSolution(queenPosition [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(queenPosition[i] + 49))
	}
	z01.PrintRune(10)
}

func EightQueens() {
	queenPosition := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	solve(0, queenPosition)
}

package main

import (
	"os"
)

func isIn(a string, r []rune) bool {
	b := []rune(a)
	if len(b) > 0 {
		for _, v := range r {
			if b[0] == v {
				return true
			}
		}
	}

	return false
}

func sum(a, b int) int {
	return a + b
}

func diff(a, b int) int {
	return a - b
}

func prod(a, b int) int {
	return a * b
}

func mod(a, b int) int {
	return a % b
}

func div(a, b int) int {
	return a / b
}

func calculator(f func(a, b int) int, c, d int) int {
	return f(c, d)
}

func main() {
	operators := []rune{'+', '-', '/', '%', '*'}
	args := os.Args[1:]
	var numA int
	var numB int
	var op string
	if len(args) >= 3 {
		numA = Atoi(args[0])
		numB = Atoi(args[2])
		op = args[1]
	}
	calc := isIn(op, operators) && IsNumeric(args[0]) && IsNumeric(args[2]) && numA < 2147483647 && numB < 2147483647 && numA > -2147483647 && numB > -2147483647
	if calc {
		if op == string(operators[0]) {
			rst := calculator(sum, numA, numB)
			PrintNbr(rst)
		}
		if op == string(operators[1]) {
			rst := calculator(diff, numA, numB)
			PrintNbr(rst)
		}
		if op == string(operators[2]) {
			if numB == 0 {
				os.Stdout.WriteString("No division by 0")
			} else {
				rst := calculator(div, numA, numB)
				PrintNbr(rst)
			}
		}
		if op == string(operators[3]) {
			if numB == 0 {
				os.Stdout.WriteString("No modulo by 0")
			} else {
				rst := calculator(mod, numA, numB)
				PrintNbr(rst)
			}
		}
		if op == string(operators[4]) {

			rst := calculator(prod, numA, numB)
			PrintNbr(rst)
		}
		os.Stdout.WriteString("\n")
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

func PrintNbr(n int) {
	issignative := n < 0
	if issignative {
		n = -1 * n
		os.Stdout.WriteString(string(rune(45)))
	} else if n == 0 {
		os.Stdout.WriteString(string(rune(48)))
	}
	slice := []uint{}
	b := uint(n)
	for b > 0 {
		slice = append(slice, b%10)
		b = b / 10
	}
	for i, j := 0, len(slice)-1; i <= j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	for i := 0; i < len(slice); i++ {
		os.Stdout.WriteString(string(rune(48 + slice[i])))
	}
}

func IsNumeric(s string) bool {
	isnumeric := false
	for i := range s {
		if i == 0 {
			if string(s[i]) == "-" || string(s[i]) == "+" {
				continue
			}
		}
		if s[i] >= '0' && s[i] <= '9' {
			isnumeric = true
		} else {
			isnumeric = false
			break
		}

	}
	return isnumeric
}

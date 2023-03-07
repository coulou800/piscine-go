package hardonexam

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// num = Atoi(nbr)
	lenBasefrom := len(baseFrom)
	lenBaseTo := len(baseTo)

	if lenBaseTo == 10 {
		rst, _ := toBase10(nbr, lenBasefrom)
		return rst
	} else {
		_, num := toBase10(nbr, lenBasefrom)
		rst := toNoDecimal(num, lenBaseTo)
		return rst
	}
}

func Atoi(s string) int {
	notnumber := false
	num := 0
	for _, v := range s {
		if v < '0' || v > '9' && v != '-' || v != '+' {
			notnumber = true
		}
	}
	if !notnumber {
		runes := []rune(s)

		for _, v := range runes {
			if v != '-' && v != '+' {
				num *= 10
				num += int(v + 48)
			}
		}
	}
	return num
}

func PowInts(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}

func toBase10(nbr string, lenBasefrom int) (string, int) {
	r := rune(48)
	for i := lenBasefrom - 1; i >= 0; i-- {
		if nbr[i] == 'f' {
			r += rune(PowInts(15, 16))
		}
		if nbr[i] == 'e' {
			r += rune(PowInts(14, 16))
		}
		if nbr[i] == 'd' {
			r += rune(PowInts(13, 16))
		}
		if nbr[i] == 'c' {
			r += rune(PowInts(12, 16))
		}
		if nbr[i] == 'b' {
			r += rune(PowInts(11, 16))
		}
		if nbr[i] == 'a' {
			r += rune(PowInts(10, 16))
		}
		r += rune(PowInts(int(rune(nbr[i]+48)), 10)) + 48
	}
	return string(r), int(r + 48)
}

func toNoDecimal(num int, lenBase int) string {
	hex := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	bin := []rune{'1', '0'}
	rst := []rune{}
	for num != 0 {
		if lenBase == 2 {
			rst = append(rst, bin[num%2])
			num /= lenBase
		} else {
			rst = append(rst, hex[num%2])
			num /= lenBase
		}
	}
	return string(rst)
}

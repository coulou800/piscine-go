package piscine

func TrimAtoi(s string) int {
	slice := []rune{}
	isneg := false
	for i := range s {
		if IsNumeric(string(s[i])) || s[i] == '+' || s[i] == '-' {
			slice = append(slice, rune(s[i]))
		}
	}
	num := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == '-' || slice[i] == '+' {
			if slice[0] == '-' {
				isneg = true
			}
			RemoveElement(slice, i)
		} else {
			num *= 10
			num += int(slice[i]) - '0'
		}
	}
	if len(slice) == 0 {
		return 0
	}
	if isneg {
		return -num
	} else if slice[0] != '-' {
		return num
	}
	return 0
}

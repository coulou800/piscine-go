package piscine

func NbrBase(nbr int, base string) string {
	baseLen := len(base)
	number := ""
	isNegative := false
	if nbr < 0 {
		isNegative = true
	}
	if nbr != 0 {
		for nbr != 0 {
			mod := nbr % baseLen
			if mod < 0 {
				mod = -mod
			}
			number += string(base[mod])
			nbr /= baseLen
		}
	} else {
		number = string(base[0])
	}
	if isNegative {
		number += "-"
	}
	number = StrRev(number)
	return number
}

func ValidateBase(base string) bool {
	baseLen := len(base)
	if baseLen <= 1 {
		return false
	} else if string(base[0]) == "-" || string(base[0]) == "+" {
		return false
	} else {
		occurrence := 0
		for i := 0; i < baseLen-1; i++ {
			occurrence = 0
			for j := 0; j < baseLen; j++ {
				if base[i] == base[j] {
					occurrence += 1
					if occurrence == 2 {
						break
					}
				}
			}
			if occurrence == 2 {
				break
			}
		}
		return occurrence != 2
	}
}

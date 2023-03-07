package piscine

func AtoiBase(s string, base string) int {
	n := StrLen(base)
	c := StrLen(s)
	if n < 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		for j := i + 1; j < n; j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	res := 0
	nb := 1
	numeric := 0
	for i := c - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if base[j] == s[i] {
				numeric = j
			}
		}
		res = res + numeric*nb
		nb *= n
	}
	return res
}

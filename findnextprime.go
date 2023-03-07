package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else if nb <= 1 {
		return 2
	}
	for {
		if IsPrime(nb) {
			break
		}
		nb++
	}
	return nb
}

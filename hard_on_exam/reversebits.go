package hardonexam

func ReverseBits(oct byte) byte {
	res := byte(0)

	for i := 0; i < 8; i++ {
		bit2 := (oct >> i) & 1
		res = res | (bit2 << (7 - i))
	}

	return res
}

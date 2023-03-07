package piscine

func Join(strs []string, sep string) string {
	rst := ""
	for i := range strs {
		if i > 0 && i < len(strs) {
			rst += sep
		}

		rst += strs[i]
	}
	return rst
}

package piscine

func SortIntegerTable(table []int) {
	for i := range table {
		for u := range table {
			if table[i] < table[u] {
				Swap(&table[i], &table[u])
			}
		}
	}
}

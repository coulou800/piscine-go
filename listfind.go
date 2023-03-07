package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	curr := l.Head
	for curr != nil {
		if comp(ref, curr.Data) {
			return &curr.Data
		}

		curr = curr.Next
	}
	return &curr.Data
}

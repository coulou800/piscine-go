package piscine

func ListSize(l *List) int {
	var c int
	if l.Head == nil {
		return 0
	}
	current := l.Head
	c++
	for current.Next != nil {
		c++
		current = current.Next
	}
	return c
}

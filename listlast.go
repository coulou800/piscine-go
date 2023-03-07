package piscine

func ListLast(l *List) interface{} {
	var n *NodeL
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for current != nil {
		n = current
		current = current.Next
	}
	return n.Data
}

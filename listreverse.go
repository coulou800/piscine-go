package piscine

func ListReverse(l *List) {
	curr := l.Head
	var prev *NodeL
	var next *NodeL

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
	l.Tail = next
}

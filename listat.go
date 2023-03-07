package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var n *NodeL
	current := l
	if pos >= 0 {
		for pos >= 0 {
			n = current
			if current != nil {
				current = current.Next
			}
			pos--
		}
		return n
	}
	return nil
}

package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head != nil {
		curr := l.Head
		var prev *NodeL
		for curr != nil {
			if curr.Data == data_ref {
				if prev == nil {
					l.Head = curr.Next
				} else {
					prev.Next = curr.Next
				}
			} else {
				prev = curr
			}
			curr = curr.Next
		}
	}
}

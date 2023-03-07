package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l

	for current != nil {
		val := l
		for val != nil {
			if current.Data < val.Data {
				temp := current.Data
				current.Data = val.Data
				val.Data = temp
			}
			val = val.Next
		}
		current = current.Next
	}
	return l
}

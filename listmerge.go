package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2.Head != nil {
		curr := l2.Head
		for curr != nil {
			ListPushBack(l1, curr.Data)
			curr = curr.Next
		}
	}
}

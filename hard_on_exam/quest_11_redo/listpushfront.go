package quest11

func ListPushFront(l *List, data interface{}) {
    n := *NodeL{Data:data}
    if head = nil {
        l.Head = n
        l.Tail = n
    } else {
        n.Next = l.Head
        l.Head = n
    }

}

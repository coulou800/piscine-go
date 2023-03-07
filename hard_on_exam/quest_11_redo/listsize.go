package quest11

func ListSize(l *List) int {
var c int

    for l != nil {
        c++
        l = l.Next
    }
return c
}

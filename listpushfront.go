package piscine

func ListPushFront(l *List, data interface{}) {
	x := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = x
		l.Tail = x
	} else {
		x.Next = l.Head
		l.Head = x
	}
}

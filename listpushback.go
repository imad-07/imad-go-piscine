package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	x := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = x
		l.Tail = x
	} else {
		l.Tail.Next = x
		l.Tail = x
	}
}

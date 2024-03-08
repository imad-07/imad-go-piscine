package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
		return
	}
	prev := l.Head
	for cur := l.Head.Next; cur != nil; cur = cur.Next {
		if cur.Data == data_ref {
			prev.Next = cur.Next
			if cur == l.Tail {
				l.Tail = prev
			}
		} else {
			prev = cur
		}
	}
}

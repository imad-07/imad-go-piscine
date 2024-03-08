package piscine

func ListSize(l *List) int {
	size := 0
	for i := l.Head; i != nil; i = i.Next {
		size++
	}
	return size
}

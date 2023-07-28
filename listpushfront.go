package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = n
	} else {
		for l.Head != nil {
			tmp := l
			l.Head = n.Next
			l.Tail = tmp.Tail
		}
	}
}

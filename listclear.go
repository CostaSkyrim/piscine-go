package piscine

func ListClear(l *List) interface{} {
	if l == nil {
		return nil
	}
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = nil
		current = next
	}
	l.Head = nil
	return l.Head
}

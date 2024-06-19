package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head

	for current != nil {
		if CompStr(current.Data, data_ref) {
			next := current.Next
			current.Next = nil
			current = next
		}
	}
}

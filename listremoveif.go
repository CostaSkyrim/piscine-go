package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	// Remove elements from the beginning of the list
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	current := l.Head

	// Traverse the list and remove matching elements
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

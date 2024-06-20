package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		// If l1 is empty, just point it to l2
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else if l2.Head != nil {
		// Otherwise, link the tail of l1 to the head of l2
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
	}
}

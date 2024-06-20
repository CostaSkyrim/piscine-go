package piscine

// SortListInsert inserts a new element into the sorted linked list
// and returns the head of the list
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	// If the list is empty or the new node should be inserted before the head
	if l == nil || l.Data >= data_ref {
		newNode.Next = l
		return newNode
	}

	// Traverse the list to find the insertion point
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert the new node
	newNode.Next = current.Next
	current.Next = newNode

	return l
}

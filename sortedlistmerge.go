package piscine

// SortedListMerge merges two sorted linked lists and returns the head of the new sorted list
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Create a dummy node to simplify the merging process
	dummy := &NodeI{}
	current := dummy

	// Merge the two lists
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Append the remaining nodes, if any
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	return dummy.Next
}

package piscine

// NodeI represents a node in the linked list with an integer data field
type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts the nodes of the linked list in ascending order using Merge Sort
func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	// Split the list into two halves
	mid := getMiddle(l)
	left := l
	right := mid.Next
	mid.Next = nil

	// Recursively sort the two halves
	left = ListSort(left)
	right = ListSort(right)

	// Merge the two sorted halves
	return merge(left, right)
}

// getMiddle finds the middle of the linked list
func getMiddle(l *NodeI) *NodeI {
	if l == nil {
		return l
	}

	slow := l
	fast := l

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// merge merges two sorted linked lists
func merge(left *NodeI, right *NodeI) *NodeI {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var result *NodeI
	if left.Data <= right.Data {
		result = left
		result.Next = merge(left.Next, right)
	} else {
		result = right
		result.Next = merge(left, right.Next)
	}

	return result
}

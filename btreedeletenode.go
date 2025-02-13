// Quest12
package piscine

// BTreeDeleteNode deletes node from the tree given by root.
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
	} else {
		successor := BTreeMin(node.Right)
		if successor.Parent != node {
			root = BTreeTransplant(root, successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
	}
	return root
}

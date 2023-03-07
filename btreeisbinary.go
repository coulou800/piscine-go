package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right != nil {
		if root.Data > root.Right.Data {
			return false
		} else {
		}
	}
	if root.Left != nil {
		if root.Data < root.Left.Data {
			return false
		}
	}

	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}

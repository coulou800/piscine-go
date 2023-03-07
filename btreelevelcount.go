package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r, l := BTreeLevelCount(root.Right), BTreeLevelCount(root.Left)
	if r < l {
		return l + 1
	} else {
		return r + 1
	}
}

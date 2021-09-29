package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 1
	}

	return BTreeLevelCount(root.Left) + BTreeLevelCount(root.Right)
}

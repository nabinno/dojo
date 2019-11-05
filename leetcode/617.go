package main

// MergeTrees ...
func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Left = MergeTrees(t1.Left, t2.Left)
	t1.Right = MergeTrees(t1.Right, t2.Right)
	t1.Val += t2.Val

	return t1
}

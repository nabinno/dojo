package main

// BstToGst ....
func BstToGst(root *TreeNode) *TreeNode {
	newRoot := new(TreeNode)
	buildNewTree(root, root, newRoot)
	return newRoot
}

func buildNewTree(root, node, newNode *TreeNode) {
	sum := 0
	dfs(root, node.Val, &sum)
	newNode.Val = sum

	if node.Left != nil {
		newNode.Left = new(TreeNode)
		buildNewTree(root, node.Left, newNode.Left)
	}
	if node.Right != nil {
		newNode.Right = new(TreeNode)
		buildNewTree(root, node.Right, newNode.Right)
	}
}

func dfs(node *TreeNode, v int, sum *int) {
	if node.Val >= v {
		*sum += node.Val
	}
	if node.Val >= v && node.Left != nil {
		dfs(node.Left, v, sum)
	}
	if node.Right != nil {
		dfs(node.Right, v, sum)
	}
}

package main

import (
	"strings"
)

func main() {
}

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// RangeSumBST (938)
func RangeSumBST(root *treeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := 0
	switch {
	case root.Val < L:
		sum = RangeSumBST(root.Right, L, R)
	case R < root.Val:
		sum = RangeSumBST(root.Left, L, R)
	default:
		sum = root.Val
		sum += RangeSumBST(root.Left, L, R)
		sum += RangeSumBST(root.Right, L, R)
	}
	return sum
}

var null = -1 << 63

func convertIntsToTreeNode(ints []int) *treeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &treeNode{
		Val: ints[0],
	}
	queue := make([]*treeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != null {
			node.Left = &treeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != null {
			node.Right = &treeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// DefangIPaddr (1108)
func DefangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// NumJewelsInStones (771)
func NumJewelsInStones(J string, S string) int {
	var gems []string

	for i := 0; i < len(S); i++ {
		if strings.Contains(J, string(S[i])) {
			gems = append(gems, string(S[i]))
		}
	}
	return len(gems)
}

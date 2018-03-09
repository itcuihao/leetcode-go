package traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var l []int
	if root == nil {
		return l
	}

	l = append(l, inorderTraversal(root.Left)...)
	l = append(l, root.Val)
	l = append(l, inorderTraversal(root.Right)...)

	return l
}

func inorderFor(root *TreeNode) []int {
	var (
		l         []int
		tempStack []*TreeNode
	)

	for root != nil || len(tempStack) > 0 {
		if root != nil {
			tempStack = append(tempStack, root)
			root = root.Left
		} else if len(tempStack) > 0 {
			root = tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]
			l = append(l, root.Val)
			root = root.Right
		}
	}
	return l
}

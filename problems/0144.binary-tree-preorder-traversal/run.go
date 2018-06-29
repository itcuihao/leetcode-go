package traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	l := make([]int, 0, 3)
	if root == nil {
		return l
	}

	l = append(l, root.Val)
	l = append(l, preorderTraversal(root.Left)...)
	l = append(l, preorderTraversal(root.Right)...)

	return l
}

// 遍历
func preorderFor(root *TreeNode) []int {
	var (
		l         []int
		tempStack []*TreeNode
	)

	for root != nil || len(tempStack) > 0 {
		l = append(l, root.Val)

		if root.Left != nil {
			if root.Right != nil {
				tempStack = append(tempStack, root.Right)
			}
			root = root.Left
		} else if root.Right != nil {
			root = root.Right
		} else {

			if len(tempStack) == 0 {
				break
			}

			root = tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]
		}
	}

	return l
}

package sum

import (
	"leetcode-go/structure"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *structure.TreeNode, sum int) bool {
	return dfs(root, sum)
}

func dfs(root *structure.TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return dfs(root.Left, sum-root.Val) || dfs(root.Right, sum-root.Val)
}

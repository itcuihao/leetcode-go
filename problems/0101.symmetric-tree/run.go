package tree

import (
	"leetcode-go/structure"
)

func isSymmetric(root *structure.TreeNode) bool {
	if root == nil {
		return false
	}
	return sym(root.Left, root.Right)
}

func sym(left, right *structure.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return sym(left.Left, right.Right) && sym(left.Right, right.Left)
}

package traversal

import "testing"

var tn = &TreeNode{
	Val: 1,
	Right: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
	},
}

func TestPreorderTraversal(t *testing.T) {
	t.Log(preorderTraversal(tn))
}
func TestPreorderFor(t *testing.T) {
	t.Log(preorderFor(tn))
}

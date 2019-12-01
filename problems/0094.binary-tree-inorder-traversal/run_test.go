package traversal

import "testing"

var tn = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 7,
		},
	},
}

func TestInorderTraversal(t *testing.T) {
	t.Log(inorderTraversal(tn))
}
func TestInorderFor(t *testing.T) {
	t.Log(inorderFor(tn))
}
func TestInorder1(t *testing.T) {
	t.Log(inorder(tn))
}

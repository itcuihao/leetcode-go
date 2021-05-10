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
			Left: &TreeNode{
				Val: 8,
			},
		},
	},
}

func TestPostorderTraversal(t *testing.T) {
	t.Log(postorderTraversal(tn))
}
func TestPostorderFor(t *testing.T) {
	t.Log(portorderFor(tn))
}

func TestPostorder(t *testing.T) {
	t.Log(postorder(tn))
}

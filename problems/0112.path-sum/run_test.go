package sum

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	tree := &structure.TreeNode{
		Val: 5,
		Left: &structure.TreeNode{
			Val: 4,
			Left: &structure.TreeNode{
				Val: 11,
				Left: &structure.TreeNode{
					Val: 7,
				},
				Right: &structure.TreeNode{
					Val: 2,
				},
			},
		},

		Right: &structure.TreeNode{
			Val: 8,
			Left: &structure.TreeNode{
				Val: 13,
			},
			Right: &structure.TreeNode{
				Val: 4,
				Right: &structure.TreeNode{
					Val: 1,
				},
			},
		},
	}
	f := hasPathSum(tree, 22)
	t.Log(f)
}

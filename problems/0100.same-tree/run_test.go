package tree

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	tree1 := &structure.TreeNode{
		Val: 1,
		Left: &structure.TreeNode{
			Val: 2,
		},

		Right: &structure.TreeNode{
			Val: 3,
		},
	}

	tree2 := &structure.TreeNode{
		Val: 1,
		Left: &structure.TreeNode{
			Val: 2,
		},

		Right: &structure.TreeNode{
			Val: 3,
		},
	}
	isSameTree(tree1, tree2)
}

package tree

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	root := structure.BinaryTree
	q := &structure.TreeNode{
		Val: 5,
	}
	p := &structure.TreeNode{
		Val: 0,
	}
	ancestor := lowestCommonAncestor(root, q, p)

	t.Log(ancestor)
}

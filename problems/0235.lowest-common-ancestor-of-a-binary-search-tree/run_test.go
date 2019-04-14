package tree

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	root := structure.BinarySearchTree
	q := &structure.TreeNode{
		Val: 5,
	}
	p := &structure.TreeNode{
		Val: 3,
	}
	// ancestor := lowestCommonAncestor(root, q, p)
	ancestor := lowestCommonAncestorStackOverflow(root, q, p)

	t.Log(ancestor)
}
func TestStackOver(t *testing.T) {
	root := structure.BinarySearchTree
	q := &structure.TreeNode{
		Val: 5,
	}
	p := &structure.TreeNode{
		Val: 3,
	}
	ancestor := lowestCommonAncestorStackOverflow(root, p, q)

	t.Log(ancestor)
}

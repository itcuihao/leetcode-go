package tree

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	tree := structure.BinaryTree101
	b := isSymmetric(tree)
	t.Log(b)
}

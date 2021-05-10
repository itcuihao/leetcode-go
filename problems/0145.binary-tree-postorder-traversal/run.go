package traversal

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	l := []int{}
	if root == nil {
		return l
	}

	l = append(l, postorderTraversal(root.Left)...)
	l = append(l, postorderTraversal(root.Right)...)
	l = append(l, root.Val)
	fmt.Println(l)
	return l
}

func portorderFor(root *TreeNode) []int {
	l := []int{}
	tempStack := []*TreeNode{}

	for root != nil || len(tempStack) > 0 {
		if root != nil {
			tempStack = append(tempStack, root)
			l = append(l, root.Val)
			root = root.Right
		} else {
			root = tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]
			root = root.Left
		}
	}
	lt := len(l)
	for i := 0; i < lt/2; i++ {
		l[i], l[lt-i-1] = l[lt-i-1], l[i]
	}
	return l
}

func postorder(root *TreeNode) []int {
	var a []int
	stk := make([]*TreeNode, 0)
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			a = append(a, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return a
}

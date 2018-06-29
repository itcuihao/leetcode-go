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
	fmt.Println(l)
	lt := len(l)
	for i := 0; i < lt/2; i++ {
		l[i], l[lt-i-1] = l[lt-i-1], l[i]
	}
	return l
}

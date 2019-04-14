package tree

import (
	"leetcode-go/structure"
)

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *structure.TreeNode) *structure.TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	switch {
	case q.Val < root.Val && p.Val < root.Val:
		return lowestCommonAncestor(root.Left, q, p)
	case q.Val > root.Val && p.Val > root.Val:
		return lowestCommonAncestor(root.Right, q, p)
	}

	return root
}

func lowestCommonAncestorStackOverflow(root, p, q *structure.TreeNode) *structure.TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}

	switch {
	case q.Val < root.Val && p.Val < root.Val:
		root = root.Left
	case q.Val > root.Val && p.Val > root.Val:
		root = root.Right
		// default: // 递归退出条件
		// 	return root
	}

	return lowestCommonAncestorStackOverflow(root, q, p)
}

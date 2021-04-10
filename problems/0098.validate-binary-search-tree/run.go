package tree

import (
	"fmt"
	"github.com/itcuihao/leetcode-go/structure"
)

func Run() {
	root := structure.BinaryTree98
	b := isValidBST(root)
	fmt.Println(b)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *structure.TreeNode) bool {
	res := traverse(root)
	if len(res) <= 1 {
		return true
	}
	tmp := res[0]
	for i := 1; i < len(res); i++ {
		if res[i] <= tmp {
			return false
		}
		tmp = res[i]
	}
	return true
}

// 中序遍历获取有序数组
func traverse(root *structure.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, traverse(root.Left)...)
	res = append(res, root.Val)
	res = append(res, traverse(root.Right)...)
	return res
}

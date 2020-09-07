package structure

import "fmt"

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	fmt.Println("start")
	for l != nil {
		fmt.Printf("%v -> ", l.Val)
		l = l.Next
	}
	fmt.Println("\nend")
}

// Lists 链表
type Lists []*ListNode

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

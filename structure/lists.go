package structure

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// Lists 链表
type Lists []*ListNode

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package cycle

import (
	"leetcode-go/structure"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *structure.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		if slow == fast {
			// if slow.Val == fast.Val { // 模拟测试
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

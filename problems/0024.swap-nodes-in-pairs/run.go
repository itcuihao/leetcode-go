package pairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度：0(n)
// 空间复杂度：0(1)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := &ListNode{}
	tmp.Next = head
	l1 := tmp
	l2 := head
	for l2 != nil && l2.Next != nil {
		cur := l2.Next.Next
		l1.Next = l2.Next
		l2.Next.Next = l2
		l2.Next = cur
		l1 = l2
		l2 = l2.Next
	}
	return tmp.Next
}

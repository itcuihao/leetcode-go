package list

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{}
	s := p
	e := p
	p.Next = head
	for i := 0; i <= n; i++ {
		s = s.Next
	}
	for s != nil {
		s = s.Next
		e = e.Next
	}
	e.Next = e.Next.Next
	return p.Next
}

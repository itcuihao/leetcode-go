package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	p := &ListNode{}

	for head != nil {
		temp := head.Next
		head.Next = p
		p = head
		head = temp
	}
	return p
}

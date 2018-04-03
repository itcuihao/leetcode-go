package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	p := new(ListNode)
	p = nil
	for head != nil {
		temp := head.Next
		head.Next = p
		p = head
		head = temp
	}
	return p
}

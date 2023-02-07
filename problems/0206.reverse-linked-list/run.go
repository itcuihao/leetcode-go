package list

import "github.com/itcuihao/leetcode-go/structure"

func reverseList(head *structure.ListNode) *structure.ListNode {

	var p *structure.ListNode
	for head != nil {
		temp := head.Next
		head.Next = p
		p = head
		head = temp
	}
	return p
}

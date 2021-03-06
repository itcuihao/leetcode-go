package list2

import (
	"github.com/itcuihao/leetcode-go/structure"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// list: 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
// out:  1 -> 2 -> 5
// tmp: 0 -> 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
//      p
// time: o(n) space: o(1)
func deleteDuplicates(head *structure.ListNode) *structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := &structure.ListNode{Val: 0}
	tmp.Next = head
	p := tmp
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			same := p.Next.Val
			for p.Next != nil && p.Next.Val == same {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return tmp.Next
}

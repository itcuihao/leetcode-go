package lists

import "leetcode-go/structure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *structure.ListNode) *structure.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	la, lb := 0, 0
	for a != nil {
		a = a.Next
		la++
	}
	for b != nil {
		b = b.Next
		lb++
	}

	if la > lb {
		abs := la - lb
		for ; abs > 0; abs-- {
			headA = headA.Next
		}
	} else {
		abs := lb - la
		for ; abs > 0; abs-- {
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil {
		// if headA.Val == headB.Val { // 如果是求相交的第一个节点，则对比值
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

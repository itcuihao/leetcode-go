package lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *ListNode

	if l1.Val > l2.Val {
		head, node = l2, l2
		l2 = l2.Next
	} else {
		head, node = l1, l1
		l1 = l1.Next
	}

	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return head
}

// PrintList print listnode
func printList(l *ListNode) {
	for l != nil {
		fmt.Println("Val:", l.Val)
		l = l.Next
	}
}

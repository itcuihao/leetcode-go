package pairs

import (
	"github.com/itcuihao/leetcode-go/structure"
)

//import "leetcode-go/structure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 时间复杂度：0(n)
// 空间复杂度：0(1)
func reverseKGroup(head *structure.ListNode, k int) *structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	count := 0
	cur := head

	for cur != nil && count != k {
		cur = cur.Next
		count++
	}

	if count == k {
		cur = reverseKGroup(cur, k)
		for count > 0 {
			count--
			tmp := head.Next
			head.Next = cur
			cur = head
			head = tmp
		}
		head = cur
	}
	return head
}

func reverseKGroup2(head *structure.ListNode, k int) *structure.ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup2(node, k)
	return newHead
}

func reverse(first *structure.ListNode, last *structure.ListNode) *structure.ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}

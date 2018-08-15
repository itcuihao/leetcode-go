package lists

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

// 分治法
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	return sort(lists, 0, len(lists)-1)
}

func sort(lists []*ListNode, lo, hi int) *ListNode {
	if lo >= hi {
		return lists[lo]
	}
	mid := (hi-lo)/2 + lo
	l1 := sort(lists, lo, mid)
	l2 := sort(lists, mid+1, hi)
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}

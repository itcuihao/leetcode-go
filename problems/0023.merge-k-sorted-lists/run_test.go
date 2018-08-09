package lists

import "testing"

func TestRun(t *testing.T) {
	var ls []*ListNode
	ls = append(ls, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	})
	ls = append(ls, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	})
	ls = append(ls, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	})
	mergeKLists(ls)
}

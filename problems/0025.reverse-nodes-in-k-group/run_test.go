package pairs

import "testing"

func TestRun(t *testing.T) {
	var ls *ListNode
	ls = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	swapPairs(ls)
}

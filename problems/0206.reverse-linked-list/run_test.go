package list

import "testing"

func TestRun(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	lr := reverseList(l)
	for lr.Next != nil {
		t.Log(lr.Val)
		lr = lr.Next
	}
}

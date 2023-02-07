package list

import (
	"testing"

	"github.com/itcuihao/leetcode-go/structure"
)

func TestRun(t *testing.T) {
	l := &structure.ListNode{
		Val: 1,
		Next: &structure.ListNode{
			Val: 2,
			Next: &structure.ListNode{
				Val: 3,
			},
		},
	}

	lr := reverseList(l)
	for lr != nil {
		t.Log(lr.Val)
		lr = lr.Next
	}
}

package list

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	n := 2

	l := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						&ListNode{},
					},
				},
			},
		},
	}
	s := removeNthFromEnd(l, n)

	for s.Next != nil {
		fmt.Println(s.Val)
		s = s.Next
	}
}

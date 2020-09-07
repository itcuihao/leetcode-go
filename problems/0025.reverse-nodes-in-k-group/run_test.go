package pairs

import (
	"github.com/itcuihao/leetcode-go/structure"
	"testing"
)

var ls = &structure.ListNode{
	Val: 1,
	Next: &structure.ListNode{
		Val: 2,
		Next: &structure.ListNode{
			Val: 3,
			Next: &structure.ListNode{
				Val: 4,
				Next: &structure.ListNode{
					Val: 5,
				},
			},
		},
	},
}

func TestRun(t *testing.T) {
	l := reverseKGroup(ls, 2)
	l.Print()
}

func TestRun2(t *testing.T) {
	l := reverseKGroup2(ls, 2)
	l.Print()
}

func TestReverse(t *testing.T) {
	node := ls
	k := 3
	for i := 0; i < k; i++ {
		if node != nil {
			node = node.Next
		}
	}
	l := reverse(ls, node)
	l.Print()
}

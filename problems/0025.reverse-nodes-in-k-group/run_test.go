package pairs

import (
	"github.com/itcuihao/leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	var ls *structure.ListNode
	ls = &structure.ListNode{
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

	reverseKGroup(ls, 2)
}

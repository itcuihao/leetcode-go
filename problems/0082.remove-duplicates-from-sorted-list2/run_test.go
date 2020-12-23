package list2

import (
	"testing"

	"github.com/itcuihao/leetcode-go/structure"
)

func TestRun(t *testing.T) {
	l := structure.ListNode82
	dl := deleteDuplicates(l)
	for dl != nil {
		t.Log(dl.Val)
		dl = dl.Next
	}
}

package list

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	l := structure.ListNode82
	dl := deleteDuplicates(l)
	for dl != nil {
		t.Log(dl.Val)
		dl = dl.Next
	}
}

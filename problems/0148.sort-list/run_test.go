package list

import (
	"fmt"
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	head := structure.ListNodeB
	l := sortListMerge(head)
	for l != nil {
		fmt.Println(l.Val)
		fmt.Println(l.Next)
		l = l.Next
	}
}

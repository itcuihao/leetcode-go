package lists

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	a := structure.ListNodeA
	b := structure.ListNodeB
	c := getIntersectionNode(b, a)
	t.Log(c)
}

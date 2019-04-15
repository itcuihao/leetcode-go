package cycle

import (
	"leetcode-go/structure"
	"testing"
)

func TestRun(t *testing.T) {
	l := structure.ListNodeCycle
	b := hasCycle(l)
	t.Log(b)
}

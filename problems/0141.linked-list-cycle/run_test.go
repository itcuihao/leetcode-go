package cycle

import (
	"testing"

	"github.com/itcuihao/leetcode-go/structure"
)

func TestRun(t *testing.T) {
	l := structure.ListNodeCycle
	b := hasCycle(l)
	t.Log(b)
}

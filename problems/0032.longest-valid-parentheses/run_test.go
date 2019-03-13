package parentheses

import (
	"testing"
)

func TestRun(t *testing.T) {

	s := "()()"

	l := longestValidParentheses(s)
	if l != 4 {
		t.Error("not")
	}
}

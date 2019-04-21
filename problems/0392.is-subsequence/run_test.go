package subsequence

import "testing"

func TestRun(t *testing.T) {
	s := "abc"
	x := "ahbgdc"
	b := isSubsequence(s, x)
	t.Log(b)
}

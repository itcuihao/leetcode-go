package matching

import "testing"

func TestIsMatching(t *testing.T) {
	s := "ab"
	t.Log(isMatch(s, "*"))
}

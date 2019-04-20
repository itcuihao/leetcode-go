package string

import "testing"

func TestRun(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"

	is := isInterleave(s1, s2, s3)
	t.Log(is)
}

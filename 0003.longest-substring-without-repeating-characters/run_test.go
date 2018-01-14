package characters

import (
	"testing"
)

// go test -v -test.run TestLengthOfLongestSubstring
func TestLengthOfLongestSubstring(t *testing.T) {
	// s := "abca"
	// s := "abcabcbb"
	s := "dvdf"
	// s := "pwwkew"
	// s := "c"
	t.Log(lengthOfLongestSubstring(s))
}

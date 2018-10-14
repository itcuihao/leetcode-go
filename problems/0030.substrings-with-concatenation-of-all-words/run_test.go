package words

import (
	"testing"
)

func TestRun(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	i := findSubstring(s, words)
	t.Log(i)
}

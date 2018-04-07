package array

import (
	"testing"
)

func TestRun(t *testing.T) {
	s := []int{1, 1, 3, 4, 4}
	t.Log(removeDuplicates(s))
}

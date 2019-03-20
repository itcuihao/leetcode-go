package array

import (
	"testing"
)

func TestRun(t *testing.T) {
	ts := []struct {
		target int
		nums   []int
		out    int
	}{
		{0, []int{5, 5, 6, 7, 0, 1, 2}, 4},
		{3, []int{5, 5, 6, 7, 0, 1, 2}, -1},
	}

	for _, v := range ts {
		if v.out == search(v.nums, v.target) {
			t.Log("success")
		} else {
			t.Log("fail")
		}
	}
}

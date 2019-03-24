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
		{0, []int{4, 5, 6, 7, 0, 1, 2}, 4},
		{3, []int{4, 5, 6, 7, 0, 1, 2}, -1},
		{2, []int{0, 1, 2}, 2},
		{7, []int{4, 5, 6, 7, 0}, 3},
	}

	for _, v := range ts {
		if v.out == search(v.nums, v.target) {
			t.Log("success")
		} else {
			t.Log("fail")
		}
	}
}

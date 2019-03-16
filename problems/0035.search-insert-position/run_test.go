package position

import (
	"testing"
)

func TestRun(t *testing.T) {

	ts := []struct {
		target int
		nums   []int
		out    int
	}{
		{5, []int{1, 3, 5, 6}, 2},
		{2, []int{1, 3, 5, 6}, 1},
		{7, []int{1, 3, 5, 6}, 4},
		{0, []int{1, 3, 5, 6}, 0},
	}

	for _, v := range ts {
		i := searchInsert(v.nums, v.target)
		if v.out != i {
			t.Logf("nums: %v, target: %d, output: %v, index: %d", v.nums, v.target, v.out, i)
		} else {
			t.Log("success")
		}
	}

}

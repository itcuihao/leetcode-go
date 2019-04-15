package array

import "testing"

func TestRun(t *testing.T) {

	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	m := findKthLargest(nums, k)
	t.Log(m)
}

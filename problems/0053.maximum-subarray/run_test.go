package subarray

import "testing"

func TestRun(t *testing.T) {
	s := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(s))
}

func TestRun2(t *testing.T) {
	s := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArrayDp(s))
}

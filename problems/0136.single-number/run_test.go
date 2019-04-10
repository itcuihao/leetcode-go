package number

import "testing"

func TestRun(t *testing.T) {
	n := []int{1, 2, 2, 4, 4}
	// s := singleNumber(n)
	// s := singleNumber1(n)
	s := singleNumber2(n)
	t.Log(s)
}

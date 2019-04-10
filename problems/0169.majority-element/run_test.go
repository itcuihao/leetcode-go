package element

import "testing"

func TestRun(t *testing.T) {
	n := []int{1, 2, 3, 4, 2, 2, 2}
	// k := majorityElement(n)
	k := majorityElement2(n)
	t.Log(k)
}

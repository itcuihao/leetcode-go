package array

import "testing"

func TestRun(t *testing.T) {
	n1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	n2 := []int{2, 5, 6}
	n := 3
	merge(n1, m, n2, n)
}

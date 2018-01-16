package arrays

import "testing"

// go test -v -test.run TestFindMedianSortedArrays
func TestFindMedianSortedArrays(t *testing.T) {
	n1 := []int{1, 3}
	n2 := []int{2, 4}
	f := findMedianSortedArrays(n1, n2)
	t.Log(f)
}

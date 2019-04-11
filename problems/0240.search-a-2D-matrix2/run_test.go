package matrix2

import "testing"

func TestRun(t *testing.T) {
	m := [][]int{
		[]int{1, 2, 3},
		[]int{3, 4, 5},
	}
	target := 4
	is := searchMatrix(m, target)
	t.Log(is)
}

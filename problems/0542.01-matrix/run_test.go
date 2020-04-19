package matrix

import "testing"

func TestRun(t *testing.T) {
	matrix := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{1, 1, 1},
	}
	a := updateMatrix(matrix)
	t.Logf("%+v", a)
}

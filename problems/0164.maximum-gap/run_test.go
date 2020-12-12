package gap

import "testing"

var (
	a = []int{3, 5, 9, 1}
)

func TestRun(t *testing.T) {
	maxSub := maximumGap(a)
	t.Log(maxSub)
}

func TestQuick(t *testing.T) {
	quickSort(a, 0, len(a)-1)
	t.Log(a)
}

func TestP(t *testing.T) {
	i := partition(a, 1, len(a)-1)
	t.Log(i)
	t.Log(a)
}

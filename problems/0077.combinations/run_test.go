package combinations

import "testing"

func TestRun(t *testing.T) {
	n := 4
	k := 2
	res := combine(n, k)
	t.Log(res)
}

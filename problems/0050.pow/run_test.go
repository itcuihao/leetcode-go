package pow

import "testing"

func TestRun(t *testing.T) {
	x := 2.00
	n := 3
	p := MyPow(x, n)
	t.Log(p)
}
func TestRun2(t *testing.T) {
	x := 2.00
	n := 2
	p := MyPow2(x, n)
	t.Log(p)
}

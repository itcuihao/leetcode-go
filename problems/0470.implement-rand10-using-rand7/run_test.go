package rand7

import "testing"

func TestRun(t *testing.T) {
	r := rand10()
	t.Log(r)
}
func TestRand7(t *testing.T) {
	r := rand7()
	t.Log(r)
}

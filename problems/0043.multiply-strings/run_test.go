package strings

import "testing"

func TestRun(t *testing.T) {
	a := "1234123473241297384743921847192"
	b := "342148720845718049587238457835729"
	c := multiply(a, b)
	t.Log(c)
}

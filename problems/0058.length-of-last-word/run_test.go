package word

import "testing"

type tcheck struct {
	input string
	out   int
}

func TestRun(t *testing.T) {
	ss := []tcheck{
		// {"Hello world", 5},
		// {"", 0},
		// {" ", 0},
		{" a", 1},
		{"a ", 1},
		// {"a b ", 1},
		// {"    ", 0},
	}
	for _, s := range ss {

		i := lengthOfLastWord(s.input)

		t.Logf("in:%v, out:%v, expect:%v", s.input, i, s.out)

	}
}

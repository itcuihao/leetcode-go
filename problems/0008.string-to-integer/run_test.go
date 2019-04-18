package integer

import "testing"

func TestMyAtoi(t *testing.T) {
	// s := "-123a..."
	// s := "    010"
	// s := "-2147483648"
	s := "2147483648"
	t.Log(myAtoi(s))
}

func TestAtoi(t *testing.T) {
	s := "-123a..."
	// s := "    010"
	// s := "-2147483648"
	// s := "2147483648"
	t.Log(atoi(s))
}

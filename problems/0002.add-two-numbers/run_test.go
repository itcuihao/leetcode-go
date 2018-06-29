package numbers

import (
	"fmt"
	"testing"
)

// go test -v -test.run TestLn
func TestLn(t *testing.T) {
	s := []int{0}
	ln := makeListNode(s)
	if ln == nil {
		t.Error("nil")
	}

	fmt.Println("ln:")

	printLN(ln)
}

func TestRun(t *testing.T) {
	s1 := []int{2, 4, 5}
	s2 := []int{5, 6, 4}

	// s1 := []int{2, 4, 3}
	// s2 := []int{5, 6, 4}

	// s1 := []int{1, 8}
	// s2 := []int{0}

	// s1 := []int{5}
	// s2 := []int{5}

	// s1 := []int{9}
	// s2 := []int{9}

	l1 := makeListNode(s1)
	l2 := makeListNode(s2)

	ln := addTwoNumbers(l1, l2)
	if ln == nil {
		t.Error("nil")
	}

	fmt.Println("ln:")

	printLN(ln)
}

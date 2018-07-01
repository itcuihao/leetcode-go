package prefix

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	strs := []string{
		"abcd",
		"bddbc",
		"dabc",
	}
	fmt.Println(longestCommonPrefix(strs))
}

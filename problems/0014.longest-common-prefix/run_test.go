package prefix

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	strs := []string{
		"abcd",
		"ab",
	}
	fmt.Println(longestCommonPrefix(strs))
}

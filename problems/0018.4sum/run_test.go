package sum

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	target := 0

	nums := []int{1, 0, -1, 0, -2, 2}
	// s := letterCombinations(target)
	s := fourSum(nums, target)
	fmt.Println(s)
}

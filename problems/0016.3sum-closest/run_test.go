package closest

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := 4
	// nums := []int{0, 0, 0, 0}
	// nums := []int{-1, 0, 1, 0}
	s := threeSumClosest(nums, target)
	fmt.Println(s)
}

package threesum

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0, 0}
	nums := []int{-1, 0, 1, 0}
	fmt.Println(threeSum(nums))
}
func TestMin(t *testing.T) {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0, 0}

	fmt.Println(min(3, 2, 1))
}

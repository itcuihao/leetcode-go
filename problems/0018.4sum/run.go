package sum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var data [][]int
	ln := len(nums)
	if ln < 4 {
		return data
	}

	sort.Ints(nums)
	return nil
}

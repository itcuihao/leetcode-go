package closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	ln := len(nums)
	if ln < 3 {
		return 0
	}

	res := nums[0] + nums[1] + nums[ln-1]
	sort.Ints(nums)

	for i := 0; i < ln-2; i++ {
		l, r := i+1, ln-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > target {
				r--
			} else {
				l++
			}

			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
		}
	}
	return res
}

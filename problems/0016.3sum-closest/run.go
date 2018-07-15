package closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	nr, delta := 0, math.MaxInt64

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)
		if l < r {
			s := nums[i] + nums[l] + nums[k]
			switch true {
			case s < target:
				i++
				if delta < target-s {
					delta = target - s
					nr = s

				}
			case s > target:
				r--
				if delta > s-target {
					delta = s - target
					nr = s
				}
			default:
				return s
			}
		}
	}
	return nr
}

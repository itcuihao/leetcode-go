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

	for i := 0; i < ln-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < ln-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			s, e := j+1, ln-1
			for s < e {
				sum := nums[i] + nums[j] + nums[s] + nums[e]

				if sum == target {
					d := []int{nums[i], nums[j], nums[s], nums[e]}

					data = append(data, d)
					for s < e && nums[s] == nums[s+1] {
						s++
					}
					for s < e && nums[e-1] == nums[e] {
						e--
					}
					s++
					e--
				} else if sum < target {
					s++
				} else {
					e--
				}
			}
		}

	}

	return data
}

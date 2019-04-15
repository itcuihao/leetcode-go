package array

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	start := 0
	end := len(nums) - 1

	for start <= end {
		i := partition(nums, start, end)

		switch {
		case i+1 < k:
			start = i + 1
		case i+1 > k:
			end = i - 1
		default:
			return nums[i]
		}
	}
	return -1
}

// 3, 2, 3, 1, 2, 4, 5, 5, 6
// K:4
func partition(nums []int, start, end int) int {
	povit := nums[start]
	s := start + 1
	e := end
	for s <= e {
		if nums[s] < povit && nums[e] > povit {
			nums[s], nums[e] = nums[e], nums[s]
			s++
			e--
		}
		if nums[s] >= povit {
			s++
		}
		if nums[e] <= povit {
			e--
		}
	}
	nums[start], nums[e] = nums[e], nums[start]
	return e
}

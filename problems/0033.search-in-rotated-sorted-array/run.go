package array

func search(nums []int, target int) int {
	ln := len(nums)
	if ln == 0 {
		return -1
	}

	n := -1

	for i := 0; i < ln; i++ {
		if nums[i] > target {
			break
		} else if nums[i] == target {
			n = i
		}
	}

	for i := ln - 1; i > 0; i-- {
		if nums[i] == target {
			n = i
		}
		if i > 0 && nums[i] < nums[i-1] {
			break
		}
	}
	return n
}

package permutation

import "fmt"

func nextPermutation(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	ln := len(nums)
	firstSmall := -1
	for i := ln - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			firstSmall = i
			break
		}
	}

	fmt.Println("firstSmall:", firstSmall)
	if firstSmall == -1 {
		reverse(nums, 0, ln-1)
		return
	}

	firstLarge := -1
	for i := ln - 1; i > firstSmall; i-- {
		if nums[i] > nums[firstSmall] {
			firstLarge = i
			break
		}
	}
	fmt.Println("firstLarge:", firstLarge)

	nums[firstSmall+1], nums[firstLarge-1] = nums[firstLarge], nums[firstSmall]
	reverse(nums, firstSmall+1, ln-1)
	// reverse(nums, firstSmall, firstLarge)
	return
}

func reverse(nums []int, l, r int) {
	for ; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

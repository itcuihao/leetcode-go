package sum

import "sort"

// 61ms
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			if v+nums[i] == target {
				return []int{k, i}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if i, ok := m[v]; ok {
			return []int{i, k}
		}
		m[target-v] = k
	}
	return nil
}

func findnum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	sort.Ints(nums)
	start := 0
	end := len(nums) - 1

	for start < end {
		cur := nums[start] + nums[end]
		switch {
		case target > cur:
			start++
		case target < cur:
			end--
		default:
			return []int{nums[start], nums[end]}
		}
	}
	return nil
}

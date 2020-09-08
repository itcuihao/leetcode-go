package subarray

import "fmt"

// 思路遍历整个数组，连续相加，索引前相加小于索引位置，则从索引位置开始相加
func maxSubArray(s []int) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return s[0]
	}

	maxSoFar, maxEndingHere := s[0], s[0]

	for i := 1; i < l; i++ {
		if maxEndingHere+s[i] > s[i] {
			maxEndingHere += s[i]
			fmt.Println("a:", maxEndingHere)
		} else {
			maxEndingHere = s[i]
			fmt.Println("b:", maxEndingHere)
		}

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		fmt.Println("c:", maxSoFar)
	}
	return maxSoFar
}

func maxSubArrayDp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

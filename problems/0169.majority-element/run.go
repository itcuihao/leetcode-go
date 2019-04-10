package element

import "sort"

func majorityElement(nums []int) int {
	c := make(map[int]int)
	for _, v := range nums {
		c[v]++
	}
	for k, v := range c {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

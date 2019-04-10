package number

import (
	"sort"
)

func singleNumber(nums []int) int {

	c := make(map[int]int)
	for _, num := range nums {
		c[num]++
	}
	for k, v := range c {
		if v == 1 {
			return k
		}
	}
	return -1
}
func singleNumber1(nums []int) int {

	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		if i+1 >= len(nums) {
			return nums[i]
		}
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
func singleNumber2(nums []int) int {

	r := 0
	for _, num := range nums {
		r ^= num
	}
	return r
}

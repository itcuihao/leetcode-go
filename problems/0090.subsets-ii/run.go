package ii

import (
	"fmt"
	"sort"
)

func Run() {
	nums := []int{1, 2, 2}
	sub := subsetsWithDup(nums)
	fmt.Println(sub)
}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	l := len(nums)
outer:
	for mask := 0; mask < 1<<l; mask++ {
		t := make([]int, 0)
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		res = append(res, t)
	}
	return res
}

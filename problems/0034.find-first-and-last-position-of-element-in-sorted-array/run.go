package array

import "fmt"

func Run() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	r := searchRange(nums, target)
	fmt.Println(r)

	nums = []int{1, 3}
	target = 1
	r = searchRange(nums, target)
	fmt.Println(r)

	nums = []int{1}
	target = 1
	r = searchRange(nums, target)
	fmt.Println(r)

	nums = []int{1, 1, 1}
	target = 1
	r = searchRange(nums, target)
	fmt.Println(r)
}

func searchRange(nums []int, target int) []int {
	l := len(nums)
	r := []int{-1, -1}
	if l == 0 {
		return r
	}

	for i := 0; i < l; i++ {
		if nums[i] == target {
			r[0] = i
			break
		}
	}
	for i := l - 1; i > r[0]; i-- {
		if nums[i] == target {
			r[1] = i
			return r
		}
	}
	if r[0] > r[1] {
		r[1] = r[0]
	}
	return r
}

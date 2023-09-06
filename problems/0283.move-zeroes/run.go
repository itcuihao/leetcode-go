package zeroes

import (
	"fmt"
)

func Run() {
	n := []int{0, 0, 0, 1}
	moveZeroes(n)
	fmt.Println(n)
}

func moveZeroes(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}
	left, right := 0, 0
	for i := 0; i < l; i++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

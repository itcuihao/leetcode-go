package ii

import "fmt"

func Run() {
	nums := []int{1, 2, 3, 1}
	sum := rob(nums)
	fmt.Println(sum)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(myRob(nums[0:len(nums)-1]), myRob(nums[1:]))
}

func myRob(nums []int) int {
	pre, cur, tmp := 0, 0, 0
	for _, num := range nums {
		tmp = cur
		cur = max(pre+num, cur)
		pre = tmp
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package immutable

import "fmt"

func Run() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	c := Constructor(nums)
	fmt.Println(c.nums)
	sub := c.SumRange(0, 2)
	fmt.Println(sub)
	sub = c.SumRange(2, 5)
	fmt.Println(sub)
	sub = c.SumRange(0, 5)
	fmt.Println(sub)
}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for ; i <= j; i++ {
		sum += this.nums[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

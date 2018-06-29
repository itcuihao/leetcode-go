package array

func removeDuplicates(nums []int) int {
	l := len(nums)

	if l <= 1 {
		return l
	}

	res := 1

	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			continue
		}

		if res!=i{
			nums[res]=nums[i]
		}

		res++
	}
	return res
}

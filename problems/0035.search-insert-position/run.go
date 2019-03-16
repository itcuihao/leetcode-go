package position

func searchInsert(nums []int, target int) int {

	// for i := 0; i < len(nums); i++ {
	// 	switch {
	// 	case target < nums[i]:
	// 		return i
	// 	case target > nums[i]:
	// 		if i == len(nums)-1 {
	// 			return i + 1
	// 		}
	// 	default:
	// 		return i
	// 	}
	// }
	// return 0

	// example:
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}

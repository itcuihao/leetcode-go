package sequence

import "fmt"

func Run() {
	//nums := []int{100, 1, 200, 2, 3}
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	l := longestConsecutive(nums)
	fmt.Println(l)
}

func longestConsecutive(nums []int) int {
	length := 0
	if len(nums) == 0 {
		return length
	}
	numsMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numsMap[num] = true
	}
	for num := range numsMap {
		if numsMap[num-1] {
			continue
		}
		curLength := 1
		curNum := num
		for numsMap[curNum+1] {
			curLength++
			curNum++
		}
		if length < curLength {
			length = curLength
		}
	}
	return length
}

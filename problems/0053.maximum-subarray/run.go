package subarray

import "fmt"

// 思路遍历整个数组，连续相加，索引前相加小于索引位置，则从索引位置开始相加
func maxSubArray(s []int) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return s[0]
	}

	maxSoFar, maxEndingHere := s[0], s[0]

	for i := 1; i < l; i++ {
		if maxEndingHere+s[i] > s[i] {
			maxEndingHere += s[i]
			fmt.Println("a:", maxEndingHere)
		} else {
			maxEndingHere = s[i]
			fmt.Println("b:", maxEndingHere)
		}

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		fmt.Println("c:", maxSoFar)
	}
	return maxSoFar
}

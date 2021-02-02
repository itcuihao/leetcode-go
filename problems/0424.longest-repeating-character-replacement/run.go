package replacement

import "fmt"

func Run() {
	s := "AABABBA"
	k := 1
	c := characterReplacement(s, k)
	fmt.Println(c)
}

func characterReplacement(s string, k int) int {
	num := make(map[uint8]int, 26)
	ls := len(s)
	maxn := 0
	left := 0
	right := 0
	for right < ls {
		indexRight := s[right] - 'A'
		num[indexRight]++
		maxn = max(maxn, num[indexRight])

		if right-left+1-maxn > k {
			indexLeft := s[left] - 'A'
			num[indexLeft]--
			left++
		}
		right++
	}
	return right - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

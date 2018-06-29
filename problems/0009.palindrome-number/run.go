package number

import (
	"strconv"
)

func isPalindrome(i int) bool {
	if i < 0 {
		return false
	}
	s := strconv.Itoa(i)
	ls := len(s)
	if ls == 1 {
		return true
	}
	for j := 0; j < ls/2; j++ {
		if s[j] != s[ls-j-1] {
			return false
		}
	}
	return true
}

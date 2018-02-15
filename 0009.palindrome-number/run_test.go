package number

import "testing"

func TestIsPalindrome(t *testing.T) {
	// i := 1221
	// i := -1221
	// i := 123123
	i := 1000021
	t.Log(isPalindrome(i))
}

package sum

import (
	"testing"
	"time"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	s1 := time.Now()
	t.Log(twoSum(nums, 9))
	t.Logf("用时:%d", time.Now().Sub(s1).Nanoseconds())

	s2 := time.Now()
	t.Log(twoSum2(nums, 13))
	t.Logf("用时:%d", time.Now().Sub(s2).Nanoseconds())
}

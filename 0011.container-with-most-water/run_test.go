package water

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{10, 10, 2}
	a := maxArea(height)
	fmt.Println(a)
	t.Logf("value:%v", a)
}

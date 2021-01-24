package water

import "fmt"

func Run() {
	height := []int{4, 2, 0, 3, 2, 5}
	t := trap(height)
	fmt.Println(t)
}

func trap(height []int) int {
	t := 0
	size := len(height)
	for i := 1; i < size-1; i++ {
		maxL, maxR := 0, 0
		for j := i; j >= 0; j-- {
			maxL = max(maxL, height[j])
		}
		for j := i; j < size; j++ {
			maxR = max(maxR, height[j])
		}
		t += min(maxL, maxR) - height[i]
	}
	return t
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

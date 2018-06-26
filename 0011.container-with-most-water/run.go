package water

func maxArea(height []int) int {
	lh := len(height)
	if lh == 0 {
		return 0
	}
	l, r := 0, lh-1

	max := 0

	minf := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	maxf := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for l < r {
		lh, rh := height[l], height[r]
		max = maxf(minf(lh, rh)*(r-l), max)

		if lh < rh {
			l++
		} else {
			r--
		}
	}
	return max
}

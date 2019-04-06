package sqrt

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	l, h := 0, x
	for l < h {
		m := l + (h-l)/2
		if x/m >= m {
			l = m + 1
		} else {
			h = m
		}
	}

	return h - 1
}

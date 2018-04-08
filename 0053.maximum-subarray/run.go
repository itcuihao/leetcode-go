package subarray

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
		} else {
			maxEndingHere = s[i]
		}

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

package integers

func divide(a int, b int) int {
	if b == 0 {
		return math.MaxInt32
	}

	if a == ^math.MaxInt32 && b == -1 {
		return ^a
	}
	ddf, dif := math.Abs(float64(a)), math.Abs(float64(b))
	dd, di := int64(ddf), int64(dif)

	ans := 0
	for dd >= di {
		tmp := 0
		for dd >= di<<uint64(tmp) {
			tmp++
		}
		t := tmp - 1

		ans += 1 << uint64(t)
		dd -= di << uint64(t)
	}

	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		return ans
	}
	return -ans
}

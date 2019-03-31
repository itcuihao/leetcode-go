package pow

import (
	"math"
)

// 递归
func MyPow(x float64, n int) float64 {

	if n > 0 {
		return pow(x, n)
	}

	return 1.0 / pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := pow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

// 第二种方法，非递归
func MyPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := float64(1)

	// 取绝对值
	absf := math.Abs(float64(n))
	abs := int64(absf)

	for abs > 0 {

		if abs%2 != 0 {
			res *= x
		}
		x *= x

		abs /= 2
	}

	if n < 0 {
		return 1.0 / res
	}
	return res
}

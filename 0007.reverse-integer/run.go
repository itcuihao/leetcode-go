package integer

import (
	"math"
	"strconv"
)

func reverse(x int) int {

	s := strconv.Itoa(x)
	i := make([]rune, len(s))
	l := len(s)
	j := 1

	if x < 0 {
		s = s[1:]
		j = -1
	}

	for k, v := range s {
		i[l-k-1] = v
	}

	for k, v := range i {
		if v == 0 {
			i = i[k+1:]
		} else {
			break
		}
	}

	s = string(i)

	x, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	if j < 0 {
		x = ^x + 1 // 按位取反 +1 得到对应负数
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	return x
}

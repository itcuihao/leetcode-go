package integer

import (
	"math"
	"strconv"
)

func reverse(x int) int {

	s := strconv.Itoa(x)
	i := make([]rune, len(s))
	l := len(s)

	if x < 0 {
		s = s[1:]
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

	j, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	if x < 0 {
		j = ^j + 1 // 按位取反 +1 得到对应负数
	}
	if j > math.MaxInt32 || j < math.MinInt32 {
		return 0
	}
	return j
}

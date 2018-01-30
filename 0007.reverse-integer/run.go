package integer

import (
	"strconv"
)

func reverse(x int) int {
	if x > 2147483647 || x < -2147483648 {
		return 0
	}
	s := strconv.Itoa(x)
	i := make([]rune, len(s))
	l := len(s)
	j := ""

	if s[0] == 45 {
		s = s[1:]
		j = "-"
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
	if j == "-" {
		x = ^x + 1 // 按位取反 +1 得到对应负数
		return x
	}
	return x
}

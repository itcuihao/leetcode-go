package integer

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)

	ls := len(str)
	if ls == 0 {
		return 0
	}

	myi, m, plus, less := 0, 1, false, false
	if str[0] == '-' {
		str = str[1:]
		less = true
	} else if str[0] == '+' {
		str = str[1:]
		plus = true
	}

	if plus || less {
		ls--
	}

	for k, v := range str {
		if v < 48 || v > 57 {
			str = str[:k]
			ls = k
			break
		}
	}
	for i := ls - 1; i >= 0; i-- {
		myi += int(str[i]-48) * m
		if isOver(myi) {
			if less {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		m *= 10
	}

	if less {
		myi = ^myi + 1
	}
	return myi
}

func isOver(i int) bool {
	if i > math.MaxInt32 {
		return true
	} else if i < math.MinInt32 {
		return true
	}
	return false
}

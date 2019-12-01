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

func atoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" || s == "0" {
		return 0
	}
	firsts := s[0]
	sign := 1
	start := 0
	res := 0
	if firsts == '+' {
		sign = 1
		start++
	} else if firsts == '-' {
		sign = -1
		start++
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return res * sign
		}
		p := int(s[i] - '0')
		res = res*10 + p
		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && res > math.MaxInt32 {
			return math.MinInt32
		}
	}
	return res * sign
}

package integer

import (
	"strconv"
)

func myAtoi(str string) int {
	var b bool
	if str[0] == '-' {
		str = str[1:]
		b = true
	}
	for _, v := range str {
		if v < 48 || v > 57 {
			return 0
		}

	}
	strconv.Atoi()
}

package say

import (
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	s := countAndSay(6)
	fmt.Println(s)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	res, start, length := make([]string, n), 0, len(s)
	for i := 1; i < length+1; i++ {
		if i == length {
			res = append(res, strconv.Itoa(i-start), string(s[start]))
			continue
		}
		if s[i] != s[start] {
			res = append(res, strconv.Itoa(i-start), string(s[start]))
			start = i
		}
	}
	return strings.Join(res, "")
}

package word

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	var last []byte
	l := len(s) - 1

	for i := l; i >= 0; i-- {
		if s[i] == 32 {
			break
		}
		last = append(last, s[i])
	}
	return len(last)
}

package words

import (
	"fmt"
	"sort"
	"strings"
)

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	var (
		a      array
		length []int
	)
	a.permutation(words, 0, len(words))
	max := len(a.eles[0])
	for i := 0; i < len(s); i++ {
		fmt.Println(s)

		for _, v := range a.eles {
			l := strings.Index(s, v)
			if l > -1 {
				length = append(length, l)
			}
		}
		if len(s) == max {
			break
		}
		s = s[max:]
		fmt.Println(length)
	}
	sort.Ints(length)
	for i := len(length) - 1; i > 0; i-- {
		if length[i] == length[i-1] {
			length = append(length[:i-1], length[i:]...)
		}
	}
	return length
}

type array struct {
	eles []string
}

func (a *array) permutation(words []string, l, h int) {

	if l == h-1 {
		a.eles = append(a.eles, strings.Join(words, ""))
	} else {
		for i := l; i < h; i++ {
			words[i], words[l] = words[l], words[i]
			a.permutation(words, l+1, h)
			words[i], words[l] = words[l], words[i]

		}
	}
}

package substring

import (
	"fmt"
	"strings"
)

// 超时
func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	// m, j := "", l-1
	// for j > 0 {
	// 	for i := 0; i < l; i++ {

	// 		fmt.Println(i, ":", j)
	// 		if i < j && s[i] == s[j] && len(s[i:j+1]) > len(m) {
	// 			fmt.Println(i, ":", j)
	// 			fmt.Println(string(s[i]), "-", string(s[j]))
	// 			fmt.Println(string(s[i:j+1]), "~", m)
	// 			if isPalindrome(string(s[i : j+1])) {
	// 				m = s[i : j+1]
	// 			}
	// 			fmt.Println("m:", m)
	// 		}
	// 	}
	// 	j--
	// }
	m, j := "", 0
	for j < l-1 {
		for i := l - 1; i > 0; i-- {
			if i > j && s[i] == s[j] && len(s[j:i+1]) > len(m) {
				max := string(s[j : i+1])
				if isPalindrome(max) {
					m = max
				}
				fmt.Println("m:", m)
			}
		}
		j++
	}
	if m == "" {
		m = string(s[j])
	}
	return m
}

// 判断是否为回文串
func isPalindrome(s string) (b bool) {
	fmt.Println("p:", s)
	l, j := len(s), len(s)-1

	for i := 0; i < l/2; i++ {
		if s[i] == s[j] {
			b = true
		} else {
			b = false
			break
		}
		j--
	}
	fmt.Println(b)
	return
}

func handleString(s string) string {
	l := len(s)
	if l == 0 {
		return "^$"
	}

	r := make([]string, 0, l)
	for _, v := range s {
		r = append(r, string(v))
	}
	return fmt.Sprintf("^#%s#$", strings.Join(r, "#"))
}

func manacher(s string) string {
	if len(s) == 0 {
		return ""
	}
	str := handleString(s)
	ls := len(str)

	p := make([]int, ls)
	c, r := 0, 0

	for i := 1; i < ls-1; i++ {
		if r > i {
			i2 := 2*c - i
			if p[i2] < r-i {
				p[i] = p[i2]
			} else {
				p[i] = p[r-i]
			}
		} else {
			p[i] = 0
		}

		for string(str[i+1+p[i]]) == string(str[i-1-p[i]]) {
			p[i]++
		}

		if i+p[i] > r {
			c = i
			r = i + p[i]
		}

	}

	maxlen, centerIndex := 0, 0
	for i := 1; i < ls-1; i++ {
		if p[i] > maxlen {
			maxlen = p[i]
			centerIndex = i
		}
	}
	fmt.Println(centerIndex, ":", maxlen)
	fmt.Println("s:", s)

	start := (centerIndex - 1 - maxlen) / 2
	end := start + maxlen

	fmt.Println("start:", start)
	fmt.Println("end:", end)

	return s[start:end]
}

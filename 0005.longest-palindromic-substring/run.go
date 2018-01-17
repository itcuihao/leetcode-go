package substring

import "fmt"

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

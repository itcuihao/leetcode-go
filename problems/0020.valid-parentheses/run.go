package parentheses

import "fmt"

func isValid(s string) bool {

	st := make([]rune, 0, len(s))

	xk := map[rune]rune{
		')': '(',
	}
	zk := map[rune]rune{
		']': '[',
	}
	dk := map[rune]rune{
		'}': '{',
	}

	for _, i := range s {
		fmt.Println(string(i))
		if len(st) == 0 {
			st = append(st, i)
			continue
		}
		if v, ok := xk[i]; ok && v == '(' {
			st = st[:len(st)-1]
			continue
		}
		if v, ok := zk[i]; ok && v == '[' {
			st = st[:len(st)-1]
			continue
		}
		if v, ok := dk[i]; ok && v == '}' {
			st = st[:len(st)-1]
			continue
		}
		st = append(st, i)
	}
	fmt.Println("st:", st)
	if len(st) == 0 {
		return true
	} else if len(st)%2 == 0 {
		fmt.Println("s:", string(st[0]), string(st[1]))
		if v, ok := xk[st[0]]; ok && v == st[1] {
			return true
		}
	}
	return false
}

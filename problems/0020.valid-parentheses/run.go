package parentheses

import "fmt"

func isValid(s string) bool {
	st := make([]rune, 0, len(s))

	check := map[string]bool{
		"()": true,
		"[]": true,
		"{}": true,
	}

	for _, i := range s {
	
		if len(st) == 0 {
			st = append(st, i)
			continue
		}

		t := string(st[len(st)-1]) + string(i)
		
		if _, ok := check[t]; ok {
			st = st[:len(st)-1]
			continue
		}
		st = append(st, i)
	}

	return len(st) == 0
}

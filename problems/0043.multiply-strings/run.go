package strings

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	digits := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i++ {
		for j := l2 - 1; j >= 0; j-- {
			product := int((num1[i] - '0') * (num2[j] - '0'))
			p1 := i + j
			p2 := i + j + 1
			sum := product + digits[p2]
			digits[p1] += sum / 10
			digits[p2] = sum % 10
		}
	}
	s := ""
	for _, v := range digits {
		if !(v == 0 && len(s) == 0) {
			s += fmt.Sprint(v)
		}
	}
	return s
}

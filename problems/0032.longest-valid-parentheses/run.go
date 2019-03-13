package parentheses

func longestValidParentheses(s string) int {
	l := len(s)
	if l <= 0 {
		return 0
	}
	res := 0
	start := -1
	q := make([]int, 0)

	for i := 0; i < l; i++ {

		if s[i] == '(' {
			q = append(q, i)
		} else {
			if len(q) == 0 {
				start = i
			} else {
				q = q[:len(q)-1]
				if len(q) == 0 {
					res = max(res, i-start)
				} else {
					res = max(res, i-q[len(q)-1])
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

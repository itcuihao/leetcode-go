package parentheses

import "fmt"

func generateParenthesis(n int) []string {

	data := []string{}
	dfs("", n, n, &data)
	fmt.Println(data)
	return data
}

func dfs(cur string, l, r int, data *[]string) {

	if l == 0 && r == 0 {
		*data = append(*data, cur)
		return
	}
	if l > 0 {
		dfs(cur+"(", l-1, r, data)
	}
	if r > 0 && l < r {
		dfs(cur+")", l, r-1, data)
	}
}

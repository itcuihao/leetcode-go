package closest

import "fmt"

// bfs
func letterCombinations(digits string) []string {
	var jp = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ld := len(digits)
	// 初始化
	r := []string{}
	if ld == 0 {
		return r
	}
	r = []string{""}
	for i := 0; i < ld; i++ {
		var tmp []string
		// 获取jp对应的下标
		index := digits[i] - '0'
		for _, k := range r {
			for _, j := range jp[index] {
				tmp = append(tmp, k+string(j))
			}
		}
		// 重新给r赋值
		r = tmp
	}
	return r
}

// dfs
func letterCombinationsDFS(digits string) []string {
	var jp = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	cur := make([]rune, len(digits))
	ans = dfs(digits, jp, 0, cur)
	return ans
}

var (
	ans []string
)

func dfs(digits string, jp []string, l int, cur []rune) []string {

	if l == len(digits) {
		if l > 0 {
			ans = append(ans, string(cur))
		}
		return ans
	}
	s := jp[digits[l]-'0']
	for _, v := range s {
		cur[l] = v
		fmt.Println("cur:", string(cur))
		ans = dfs(digits, jp, l+1, cur)
	}
	return ans
}

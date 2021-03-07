package partitioning

import "fmt"

func Run() {
	s := "aab"
	p := partition(s)
	fmt.Println(p)
}

func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	ans := make([][]string, 0)
	splists := make([]string, 0)
	dfs := func(int) {}
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, splists...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splists = append(splists, s[i:j+1])
				dfs(j + 1)
				splists = splists[:len(splists)-1]
			}
		}
	}
	dfs(0)
	return ans
}

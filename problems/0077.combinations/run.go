package combinations

func combine(n int, k int) [][]int {

	tmp := make([]int, k)
	var res [][]int

	var dfs func(int, int)
	dfs = func(index, begin int) {
		if index == k {
			cur := make([]int, k)
			copy(cur, tmp)
			res = append(res, cur)
			return
		}
		for i := begin; i <= n+1-k+index; i++ {
			tmp[index] = i
			dfs(index+1, i+1)
		}
	}

	dfs(0, 1)
	return res
}

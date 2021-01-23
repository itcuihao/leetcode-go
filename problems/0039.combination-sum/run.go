package sum

import "fmt"

func Run() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	sums := combinationSum(candidates, target)
	fmt.Println(sums)
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	dfs := func(target, index int) {}
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), comb...))
			return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return res
}

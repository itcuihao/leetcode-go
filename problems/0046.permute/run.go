package permute

import "fmt"

func Run() {
	arr := []int{1, 2, 3}
	res := permute(arr)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return nil
	}
	res := make([][]int, 0)
	used := make([]bool, l)
	path := make([]int, l)
	df := func(int) {}
	df = func(i int) {
		if i == l {
			res = append(res, append([]int(nil), path...))
			return
		}
		for k, find := range used {
			if find {
				continue
			}
			path[i] = nums[k]
			used[k] = true
			df(i + 1)
			used[k] = false
		}
	}
	df(0)
	return res
}

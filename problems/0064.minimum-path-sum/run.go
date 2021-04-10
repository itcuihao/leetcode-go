package sum

import "fmt"

func Run() {
	grid := [][]int{{1, 3, 1}, {1, 50, 1}, {40, 2, 1}}
	m := minPathSum(grid)
	fmt.Println(m)
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	r := len(grid)
	c := len(grid[0])
	min := func(m, n int) int {
		if m > n {
			return n
		}
		return m
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == j && i == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[r-1][c-1]
}

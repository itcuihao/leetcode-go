package matrix

import (
	"container/list"
	"fmt"
)

// TODO: 先找个合适的队列
func updateMatrix(matrix [][]int) [][]int {
	list.New()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])

		}
	}
	return nil
}

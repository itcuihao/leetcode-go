package matrix

import "fmt"

func updateMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])

		}
	}
	return nil
}

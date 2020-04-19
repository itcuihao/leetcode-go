package matrix

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	maxRow, maxCel := len(matrix), len(matrix[0])
	zero := make([][]int, 0)
	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCel; j++ {
			if matrix[i][j] == 0 {
				zero = append(zero, []int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}

	around := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(zero) > 0 {
		item := zero[0]
		x := item[0]
		y := item[1]

		zero = zero[1:]

		for _, a := range around {
			ax, ay := x+a[0], y+a[1]

			if ax < 0 || ax >= maxRow || ay < 0 || ay >= maxCel {
				continue
			}

			if matrix[ax][ay] == -1 {
				matrix[ax][ay] = matrix[x][y] + 1
				zero = append(zero, []int{ax, ay})
			}

		}
	}
	return matrix
}

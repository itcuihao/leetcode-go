package matrix2

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	r, c := 0, n-1
	for r < m && c >= 0 {
		switch {
		case matrix[r][c] > target:
			c--
		case matrix[r][c] < target:
			r++
		case matrix[r][c] == target:
			return true
		}
	}

	return false
}

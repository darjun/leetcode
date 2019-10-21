package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}

	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	seen := make([][]bool, m, m)
	for i := range seen {
		seen[i] = make([]bool, n, n)
	}

	dh := []int{1, 0, -1, 0}
	dv := []int{0, 1, 0, -1}

	ans := make([]int, m*n, m*n)
	var row, col int
	var di int
	for i := 0; i < m*n; i++ {
		ans[i] = matrix[row][col]
		seen[row][col] = true

		nrow := row + dv[di]
		ncol := col + dh[di]
		if 0 <= nrow && nrow < m && 0 <= ncol && ncol < n && !seen[nrow][ncol] {
			row = nrow
			col = ncol
		} else {
			di = (di + 1) % 4
			row += dv[di]
			col += dh[di]
		}
	}

	return ans
}

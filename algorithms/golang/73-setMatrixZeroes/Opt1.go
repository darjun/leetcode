package main

func setZeroesOpt1(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	if col == 0 {
		return
	}

	firstColumn := false
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			firstColumn = true
		}

		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}

	if firstColumn {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}

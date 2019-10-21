package main

func setZeroes(matrix [][]int) {
	// value = 1:row value = 2: col
	mapRowCols := make(map[int]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				mapRowCols[i] |= 1
				mapRowCols[j] |= 2
			}
		}
	}

	for i, value := range mapRowCols {
		if value&1 != 0 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}

		if value&2 != 0 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
}

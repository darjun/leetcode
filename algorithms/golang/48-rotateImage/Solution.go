package main

func rotate(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	/*
		先反转
		1 2 3
		4 5 6
		7 8 9

		----->

		1 4 7
		2 5 8
		3 6 9
	*/
	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	/*
		交换列

		1 4 7
		2 5 8
		3 6 9

		----->

		7 4 1
		8 5 2
		9 6 3
	*/

	for i := 0; i < m; i++ {
		for j := 0; j < m/2; j++ {
			matrix[i][j], matrix[i][m-j-1] = matrix[i][m-j-1], matrix[i][j]
		}
	}
}

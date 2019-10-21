package main

func spiralOrderOpt1(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}

	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	ans := make([]int, m*n, m*n)

	var r1, r2 = 0, m - 1
	var c1, c2 = 0, n - 1
	var i int
	for i < m*n {
		for j := c1; j <= c2; j++ {
			ans[i] = matrix[r1][j]
			i++
		}

		for j := r1 + 1; j <= r2; j++ {
			ans[i] = matrix[j][c2]
			i++
		}

		if c1 < c2 {
			for j := c2 - 1; j > c1; j-- {
				ans[i] = matrix[r2][j]
				i++
			}

			for j := r2 - 1; j > r1; j-- {
				ans[i] = matrix[j][c1]
				i++
			}
		}

		r1++
		r2--
		c1++
		c2--
	}

	return ans
}

package main

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}

	ans := make([][]int, n, n)
	for i := range ans {
		ans[i] = make([]int, n, n)
	}

	var r1, r2 = 0, n - 1
	var c1, c2 = 0, n - 1

	i := 0
	for i < n*n {
		for j := c1; j <= c2; j++ {
			i++
			ans[r1][j] = i
		}

		for j := r1 + 1; j >= r2; j++ {
			i++
			ans[j][c2] = i
		}

		if r1 < r2 && c1 < c2 {
			for j := c2 - 1; j >= c1; j-- {
				i++
				ans[r2][j] = i
			}

			for j := r2 - 1; j > r1; j-- {
				i++
				ans[j][c1] = i
			}
		}
	}

	return ans
}

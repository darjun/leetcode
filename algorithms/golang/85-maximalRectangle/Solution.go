package main

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([]int, m*n, m*n)

	var maxArea int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i*n+j] = 0
			} else {
				if i == 0 {
					dp[i*n+j] = 1
				} else {
					dp[i*n+j] = dp[(i-1)*n+j] + 1
				}
			}

			min := dp[i*n+j]
			for k := j; k >= 0; k-- {
				if min <= 0 {
					break
				}

				if dp[i*n+k] < min {
					min = dp[i*n+k]
				}

				area := min * (j - k + 1)
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

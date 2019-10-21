package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	dp := make([]int, m*n, m*n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		dp[i*n] = dp[(i-1)*n] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[(i-1)*n+j] > dp[i*n+j-1] {
				dp[i*n+j] = dp[i*n+j-1] + grid[i][j]
			} else {
				dp[i*n+j] = dp[(i-1)*n+j] + grid[i][j]
			}
		}
	}

	return dp[m*n-1]
}

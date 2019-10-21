package main

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([]int, m*n, m*n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		dp[i*n] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i*n+j] = dp[(i-1)*n+j] + dp[i*n+j-1]
		}
	}

	return dp[m*n-1]
}

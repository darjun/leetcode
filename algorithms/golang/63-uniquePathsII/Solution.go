package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([]int, m*n, m*n)
	hasObstacle := false
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			hasObstacle = true
		}
		if hasObstacle {
			dp[i] = 0
		} else {
			dp[i] = 1
		}
	}

	hasObstacle = false
	for j := 1; j < m; j++ {
		if obstacleGrid[j][0] == 1 {
			hasObstacle = true
		}
		if hasObstacle {
			dp[j*n] = 0
		} else {
			dp[j*n] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i*n+j] = 0
			} else {
				dp[i*n+j] = dp[(i-1)*n+j] + dp[i*n+j-1]
			}
		}
	}

	return dp[m*n-1]
}

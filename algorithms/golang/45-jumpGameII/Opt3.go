package main

func jumpOpt3(nums []int) int {
	l := len(nums)
	dp := make([]int, l, l)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < l-1; i++ {
		for j := 1; j <= nums[i] && i+j < l; j++ {
			if dp[i+j] == -1 || dp[i+j] > dp[i]+1 {
				dp[i+j] = dp[i] + 1
			}
		}
	}

	return dp[l-1]
}

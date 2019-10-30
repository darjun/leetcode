package main

func singleProfit(gap []int, start, end int) int {
	n := end - start
	sum := gap[start]
	max := sum

	for i := 1; i < n; i++ {
		sum = sum + gap[i+start]
		if sum < gap[i+start] {
			sum = gap[i+start]
		}
		if sum > max {
			max = sum
		}
	}

	if max < 0 {
		return 0
	}

	return max
}

func maxProfitOpt1(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	gap := make([]int, n-1, n-1)
	for i := 0; i < n-1; i++ {
		gap[i] = prices[i+1] - prices[i]
	}

	ans := singleProfit(gap, 0, n-1)
	if n <= 2 {
		return ans
	}

	for i := 1; i < n-1; i++ {
		if gap[i] > 0 {
			continue
		}

		if gap[i] <= 0 && gap[i-1] <= 0 {
			continue
		}

		s := singleProfit(gap, 0, i) + singleProfit(gap, i, n-1)
		if s > ans {
			ans = s
		}
	}

	return ans
}

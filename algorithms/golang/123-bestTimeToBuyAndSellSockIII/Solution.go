package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	b1 := make([]int, len(prices), len(prices))
	b2 := make([]int, len(prices), len(prices))

	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		b1[i] = prices[i] - min
		if b1[i] < b1[i-1] {
			b1[i] = b1[i-1]
		}
	}

	max := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		b2[i] = max - prices[i]
		if b2[i] < b2[i+1] {
			b2[i] = b2[i+1]
		}
	}

	var ans int

	for i := 0; i < len(prices); i++ {
		s := b1[i] + b2[i]
		if s > ans {
			ans = s
		}
	}

	return ans
}

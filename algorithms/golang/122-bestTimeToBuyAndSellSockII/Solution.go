package main

func maxProfit(prices []int) int {
	var max int

	l := len(prices)
	for i := 0; i < l-1; i++ {
		if prices[i+1] > prices[i] {
			max += prices[i+1] - prices[i]
		}
	}

	return max
}

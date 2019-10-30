package main

func maxProfit(prices []int) int {
	min := -1
	var max int

	for _, price := range prices {
		if min == -1 || price < min {
			min = price
		} else if price-min > max {
			max = price - min
		}
	}

	return max
}

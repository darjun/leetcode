package main

import "math"

func maximumProductOpt1(nums []int) int {
	min1 := int(math.MaxInt32)
	min2 := int(math.MaxInt32)
	max1 := int(math.MinInt32)
	max2 := int(math.MinInt32)
	max3 := int(math.MinInt32)

	for _, num := range nums {
		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num <= min2 {
			min2 = num
		}

		if num >= max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max3 = max2
			max2 = num
		} else if num >= max3 {
			max3 = num
		}
	}

	maxProduct := max1 * max2 * max3
	product := min1 * min2 * max1
	if maxProduct < product {
		maxProduct = product
	}

	return maxProduct
}

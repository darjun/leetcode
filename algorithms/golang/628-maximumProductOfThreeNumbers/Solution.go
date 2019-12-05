package main

import "sort"

func maximumProduct(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)

	maxProduct := nums[0] * nums[1] * nums[n-1]
	product := nums[n-1] * nums[n-2] * nums[n-3]
	if maxProduct < product {
		maxProduct = product
	}

	return maxProduct
}

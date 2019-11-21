package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	
	n := len(nums)
	if n & 1 != 0 {
		return 0
	}

	var sum int
	for i := 0; i < n; i+=2 {
		sum += nums[i]
	}

	return sum
}
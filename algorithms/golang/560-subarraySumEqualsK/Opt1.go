package main

func subarraySumOpt1(nums []int, k int) int {
	var count int

	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}

	return count
}
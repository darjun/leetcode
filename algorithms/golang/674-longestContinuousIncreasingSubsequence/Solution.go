package main

func findLengthOfLCIS(nums []int) int {
	var max int
	var count int

	for i, num := range nums {
		if i == 0 || num > nums[i-1] {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 1
		}
	}

	return max
}

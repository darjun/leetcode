package main

func maxSubArrayOpt1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	curSum := nums[0]

	for _, num := range nums[1:] {
		if curSum < 0 {
			curSum = 0
		}

		curSum += num

		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum
}

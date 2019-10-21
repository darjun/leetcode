package main

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	sums := make([]int, l, l)
	sums[0] = nums[0]
	maxSum := sums[0]
	for i := 1; i < l; i++ {
		sums[i] = sums[i-1] + nums[i]
		if maxSum < sums[i] {
			maxSum = sums[i]
		}
	}

	for i := 1; i < l; i++ {
		for j := i; j < l; j++ {
			sumij := sums[j] - sums[i-1]
			if maxSum < sumij {
				maxSum = sumij
			}
		}
	}

	return maxSum
}

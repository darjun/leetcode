package main

func findMaxAverageOpt1(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	max := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if max < sum {
			max = sum
		}
	}

	return float64(max) / float64(k)
}

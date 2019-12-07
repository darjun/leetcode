package main

func findMaxAverage(nums []int, k int) float64 {
	sums := make([]int, len(nums), len(nums))
	sums[0] = nums[0]
	for i, num := range nums[1:] {
		sums[i+1] = sums[i] + num
	}

	max := float64(sums[k-1]) / float64(k)
	for i := k; i < len(nums); i++ {
		average := float64(sums[i]-sums[i-k]) / float64(k)
		if max < average {
			max = average
		}
	}

	return max
}

package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastValue := nums[0]
	count := 1
	for _, value := range nums[1:] {
		if lastValue != value {
			lastValue = value
			nums[count] = value
			count++
		}
	}

	return count
}

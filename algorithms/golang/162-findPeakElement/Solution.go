package main

func findPeakElement(nums []int) int {
	index := 0
	l := len(nums)

	for index < l-1 && nums[index] < nums[index+1] {
		index++
	}

	if index == 0 || nums[index] > nums[index-1] {
		return index
	}

	return -1
}

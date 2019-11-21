package main

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	low := 0
	high := len(nums)-1

	for low < high && nums[low] <= nums[low+1] {
		low++
	}

	if low == high {
		return 0
	}

	for low < high && nums[high] >= nums[high-1] {
		high--
	}

	min := int(^uint(0) >> 1)
	max := ^min

	for i := low; i <= high; i++ {
		if nums[i] < min {
			min = nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	low--
	for ; low >= 0 && nums[low] > min; low-- {
	}

	high++
	for ; high < n && nums[high] < max; high++ {
	}

	return high - low - 1
}
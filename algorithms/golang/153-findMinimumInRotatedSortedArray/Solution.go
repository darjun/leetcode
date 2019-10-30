package main

func findMin(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return nums[lo]
}

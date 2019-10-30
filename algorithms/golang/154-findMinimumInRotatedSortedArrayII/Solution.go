package main

func findMin(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	l := len(nums)
	for lo < hi && nums[hi] == nums[hi-1] {
		hi--
	}

	for lo < hi && nums[lo] == nums[l-1] {
		lo++
	}

	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] <= nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return nums[lo]
}

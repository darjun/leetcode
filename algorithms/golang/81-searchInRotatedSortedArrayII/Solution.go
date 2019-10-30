package main

func search(nums []int, target int) bool {
	lo := 0
	hi := len(nums) - 1
	lastIndex := len(nums) - 1
	for hi > 0 && nums[hi] == nums[hi-1] {
		hi--
	}

	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] < nums[lastIndex] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	if lo == 0 {
		hi = lastIndex
	} else {
		if target >= nums[0] && target <= nums[lo-1] {
			hi = lo - 1
			lo = 0
		} else if target >= nums[lo] && target <= nums[lastIndex] {
			hi = lastIndex
		} else {
			return false
		}
	}

	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return true
}

package main

func searchRange(nums []int, target int) []int {
	lo := 0
	hi := len(nums) - 1
	left := -1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			left = mid
		}

		if nums[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if left == -1 {
		return []int{-1, -1}
	}

	lo = left
	hi = len(nums) - 1
	right := left
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			lo = mid + 1
			right = mid
		} else {
			hi = mid - 1
		}
	}

	return []int{left, right}
}

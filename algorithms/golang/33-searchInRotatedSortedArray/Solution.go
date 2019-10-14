package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	n := len(nums)
	lo := 0
	hi := n - 1

	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if lo == 0 {
		// 没有反转
		hi = n - 1
	} else {
		if target <= nums[n-1] {
			// target只能在第二部分,lo已经最小值的索引了
			hi = n - 1
		} else {
			// target在第一部分,lo-1是最大值的索引
			hi = lo - 1
			lo = 0
		}
	}

	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

package main

import "sort"

func binarySearch(nums []int, l, r, target int) int {
	for l <= r && r < len(nums) {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func triangleNumberOpt1(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var count int
	for i := 0; i < n-2; i++ {
		if nums[i] == 0 {
			continue
		}

		k := i + 2
		for j := i + 1; j < n-1; j++ {
			k = binarySearch(nums, k, n-1, nums[i]+nums[j])
			count += k - j - 1
		}
	}

	return count
}

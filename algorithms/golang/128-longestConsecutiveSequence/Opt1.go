package main

import "sort"

func longestConsecutiveOpt1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	var maxLen int
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				curLen++
			} else {
				if maxLen < curLen {
					maxLen = curLen
				}
				curLen = 1
			}
		}
	}

	if maxLen < curLen {
		maxLen = curLen
	}

	return maxLen
}

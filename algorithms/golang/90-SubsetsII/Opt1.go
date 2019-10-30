package main

import "sort"

func helperOpt1(nums []int, index int, subset []int, ans *[][]int) {
	if index > len(nums) {
		return
	}

	*ans = append(*ans, cloneSlice(subset))
	for i := index; i < len(nums); i++ {
		if i > index && nums[i-1] == nums[i] {
			continue
		}

		subset = append(subset, nums[i])
		helperOpt1(nums, i+1, subset, ans)
		subset = subset[:len(subset)-1]
	}
}

func subsetsWithDupOpt1(nums []int) [][]int {
	ans := make([][]int, 0, len(nums))

	sort.Ints(nums)

	helperOpt1(nums, 0, make([]int, 0, 1), &ans)
	return ans
}

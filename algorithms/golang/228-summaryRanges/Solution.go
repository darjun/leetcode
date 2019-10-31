package main

import "strconv"

func summaryRanges(nums []int) []string {
	var ans []string
	if len(nums) == 0 {
		return ans
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[j-1]+1 {
			if i == j-1 {
				ans = append(ans, strconv.FormatInt(int64(nums[i]), 10))
			} else {
				ans = append(ans, strconv.FormatInt(int64(nums[i]), 10)+"->"+strconv.FormatInt(int64(nums[j-1]), 10))
			}
			i = j
		}
	}

	if i == len(nums)-1 {
		ans = append(ans, strconv.FormatInt(int64(nums[i]), 10))
	} else {
		ans = append(ans, strconv.FormatInt(int64(nums[i]), 10)+"->"+strconv.FormatInt(int64(nums[len(nums)-1]), 10))
	}

	return ans
}

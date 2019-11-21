package main

func arrayNestingOpt2(nums []int) int {
	var maxLen int
	for i := 0; i < len(nums); i++ {
		var count int
		for k := i; nums[k] >= 0; {
			ak := nums[k]
			nums[k] = -nums[k]-1
			count++
			k = ak;
		}

		if maxLen < count {
			maxLen = count
		}
	}

	return maxLen
}
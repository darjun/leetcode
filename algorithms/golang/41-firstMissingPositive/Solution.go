package main

func firstMissingPositive(nums []int) int {
	i := 0
	l := len(nums)
	for ; i < l; i++ {
		j := nums[i]
		if j <= 0 || j > l {
			continue
		}

		if nums[i] != nums[j-1] {
			nums[i], nums[j-1] = nums[j-1], nums[i]
			i--
		}
	}

	for i = 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return l + 1
}

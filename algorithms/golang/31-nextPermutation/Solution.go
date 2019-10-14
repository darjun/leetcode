package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	var start, end int
	if i == -1 {
		start = 0
		end = len(nums) - 1
	} else {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		start = i + 1
		end = len(nums) - 1
	}

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

package main

func arrayNesting(nums []int) int {
	var maxLen int
	for i := 0; i < len(nums); i++ {
		count := 1
		slow := nums[i]
        fast := nums[slow]
		for slow != fast {
			slow = nums[slow]
			fast = nums[nums[fast]]
			count++
		}

		if maxLen < count {
			maxLen = count
		}
	}

	return maxLen
}
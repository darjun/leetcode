package main

func checkPossibility(nums []int) bool {
	index := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if index != -1 {
				return false
			}

			index = i
		}
	}

	if index == -1 || index == 0 || index == len(nums) - 2 {
		return true
	}

	if nums[index+2] > nums[index] {
		return true
	}

	if nums[index+1] > nums[index-1] {
		return true
	}

	return false
}

package main

func majorityElement(nums []int) int {
	index := -1
	count := 0
	for i, num := range nums {
		if index == -1 || nums[index] == num {
			count++
			index = i
		} else {
			count--
			if count == 0 {
				index = -1
			}
		}
	}

	return nums[index]
}

package main

func removeElement(nums []int, val int) int {
	count := 0
	for _, value := range nums {
		if val != value {
			nums[count] = value
			count++
		}
	}

	return count
}

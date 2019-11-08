package main

func findDuplicates(nums []int) []int {
	var ans []int
	for _, num := range nums {
		var index int
		if num < 0 {
			index = -num -1
		} else {
			index = num - 1
		}
		if nums[index] < 0 {
			ans = append(ans, index+1)
		}
		nums[index] = -nums[index]
	}

	return ans
}
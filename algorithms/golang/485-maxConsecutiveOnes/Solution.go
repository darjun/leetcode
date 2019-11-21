package main

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var cur int
	for _, num := range nums {
		if num == 1 {
			cur++
		} else {
		 	if cur > max {
				max = cur
			}
			cur = 0
		}
	}

	if cur > max {
		max = cur
	}
	return max
}
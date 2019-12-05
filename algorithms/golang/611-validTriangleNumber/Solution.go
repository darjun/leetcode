package main

import "sort"

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var count int
	for i := 0; i < n-2; i++ {
		a := nums[i]
		for j := i + 1; j < n-1; j++ {
			b := nums[j]
			max := a + b

			for k := j + 1; k < n; k++ {
				c := nums[k]
				if c >= max {
					break
				}

				count++
			}
		}
	}

	return count
}

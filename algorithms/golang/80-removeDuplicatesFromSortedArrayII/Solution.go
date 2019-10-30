package main

func removeDuplicates(nums []int) int {
	var count int
	var i int
	var current int
	for _, num := range nums {
		needAppend := true
		if count == 0 || current == num {
			count++
			if count > 2 {
				needAppend = false
			}
		} else {
			count = 1
		}

		if needAppend {
			current = num
			nums[i] = num
			i++
		}
	}

	return i
}

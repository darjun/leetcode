package main

func moveZeroes(nums []int) {
	var index int
	for _, num := range nums {
		if num != 0 {
			nums[index] = num
			index++
		}
	}

	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

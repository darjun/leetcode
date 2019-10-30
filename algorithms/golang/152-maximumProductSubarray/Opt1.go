package main

func maxInts(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func minInts(nums []int) int {
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}

func maxProductOpt1(nums []int) int {
	max := make([]int, len(nums), len(nums))
	min := make([]int, len(nums), len(nums))

	max[0] = nums[0]
	min[0] = nums[0]
	ans := nums[0]

	for i, num := range nums[1:] {
		v1 := max[i] * num
		v2 := min[i] * num

		max[i+1] = maxInts([]int{v1, v2, num})
		min[i+1] = minInts([]int{v1, v2, num})

		if ans < max[i+1] {
			ans = max[i+1]
		}
	}

	return ans
}

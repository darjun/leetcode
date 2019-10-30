package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var ans int
	var max, min int

	for _, num := range nums {
		if num == 0 {
			max = 0
			min = 0
		} else if num > 0 {
			if max == 0 {
				max = num
			} else {
				max = max * num
			}

			if min != 0 {
				min = min * num
			}
		} else {
			if max == 0 {
				max = min * num
				min = num
			} else {
				tmp := min
				min = max * num
				max = tmp * num
			}
		}

		if max > ans {
			ans = max
		}
	}

	return ans
}

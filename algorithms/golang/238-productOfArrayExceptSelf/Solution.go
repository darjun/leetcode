package main

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums), len(nums))

	zeroIndex := -1
	product := 1
	for i, num := range nums {
		if num == 0 {
			if zeroIndex != -1 {
				// two zero
				return ans
			}

			zeroIndex = i
		} else {
			product *= num
		}
	}

	if zeroIndex != -1 {
		ans[zeroIndex] = product
		return ans
	}

	for i, num := range nums {
		ans[i] = product / num
	}

	return ans
}

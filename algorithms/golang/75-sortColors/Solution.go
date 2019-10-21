package main

func sortColors(nums []int) {
	var i, j, k int

	// values in range [0, i] = 0
	// values in range [i+1, k] = 1
	// values in range [j, len(nums)-1] = 2

	for j = len(nums) - 1; i < j; j-- {
		if nums[j] != 2 {
			j++
			break
		}
	}
	if j == i {
		return
	}

	for i = 0; i < j; i++ {
		if nums[i] != 0 {
			i--
			break
		}
	}

	if i == j {
		return
	}

	for k = i + 1; k < j; {
		for ; k < j; k++ {
			if nums[k] == 0 {
				i++
				nums[i], nums[k] = nums[k], nums[i]
			} else if nums[k] == 2 {
				j--
				nums[j], nums[k] = nums[k], nums[j]
				k--
			}
		}
	}
}

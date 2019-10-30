package main

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 || k == 0 {
		return
	}

	k %= l

	i := 0
	j := l - k - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	i = l - k
	j = l - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	i = 0
	j = l - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

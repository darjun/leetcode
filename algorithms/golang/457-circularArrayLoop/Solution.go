package main

func getIndex(i int, nums []int) int {
	n := len(nums)

	index := (i + nums[i]) % n
	if index >= 0 {
		return index % n
	}

	return n+index;
}

func circularArrayLoop(nums []int) bool {
	for i, num := range nums {
		if num == 0 {
			continue
		}

		slow, fast := i, getIndex(i, nums)
		for nums[slow] * nums[fast] > 0 && nums[getIndex(fast, nums)] * nums[i] > 0 {
			if slow == fast {
				if slow == getIndex(slow, nums) {
					break
				}
				return true
			}
			slow = getIndex(slow, nums)
			fast = getIndex(getIndex(fast, nums), nums)
		}
		slow = i
		val := nums[i]
		for nums[slow] * val > 0 {
			next := getIndex(slow, nums)
			nums[slow] = 0
			slow = next
		}

	}
	return false
}
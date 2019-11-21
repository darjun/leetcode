package main

func arrayNestingOpt1(nums []int) int {
	var maxLen int
	for i := 0; i < len(nums); i++ {
		slow := nums[i]
		if slow < 0 {
			continue
		}
        nums[i] = -nums[i]

		fast := nums[-nums[i]]
		if fast > 0 {
			nums[-nums[i]] = -fast
        } else {
            fast = -fast
        }

		count := 1
		for slow != fast {
			if nums[slow] > 0 {
				nums[slow] = -nums[slow]
			}
			slow = -nums[slow]

			if nums[fast] > 0 {
				nums[fast] = -nums[fast]
			}
			fast = -nums[fast]
			if nums[fast] > 0 {
				nums[fast] = -nums[fast]
			}
			fast = -nums[fast]
			count++
		}

		if maxLen < count {
			maxLen = count
		}
	}

	return maxLen
}
package main

func canJump(nums []int) bool {
	var nextJump int
	l := len(nums)
	for i := 0; i < l; {
		if nums[i] >= l-i-1 {
			return true
		}

		var maxReach int
		for j := i + 1; j <= i+nums[i]; j++ {
			if j+nums[j] > maxReach {
				maxReach = j + nums[j]
				nextJump = j
			}
		}

		if i == nextJump {
			return false
		} else {
			i = nextJump
		}
	}

	return false
}

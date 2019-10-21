package main

func jumpOpt2(nums []int) int {
	var jumps, maxReach, nextJump int

	l := len(nums)
	for i := 0; i < l-1; {
		if nums[i] >= l-i-1 {
			jumps++
			break
		}

		for j := i + 1; j <= i+nums[i]; j++ {
			if j+nums[j] > maxReach {
				maxReach = j + nums[j]
				nextJump = j
			}
		}

		if i == nextJump {
			jumps = -1
			break
		} else {
			i = nextJump
			jumps++
		}
	}

	return jumps
}

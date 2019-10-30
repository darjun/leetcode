package main

func minSubArrayLen(s int, nums []int) int {
	sum := 0
	var start int
	var minLen int
	for i, num := range nums {
		if num >= s {
			return 1
		}

		sum += num
		if sum >= s {
			for start < i {
				if sum-nums[start] < s {
					break
				}

				sum -= nums[start]
				start++
			}

			if minLen == 0 || minLen > i-start+1 {
				minLen = i - start + 1
			}
		}
	}

	return minLen
}

package main

func longestConsecutiveOpt2(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		m[num] = struct{}{}
	}

	var maxLen int
	for _, num := range nums {
		if _, exist := m[num-1]; exist {
			continue
		}

		curLen := 1
		for {
			num++
			_, exist := m[num]
			if !exist {
				break
			}

			curLen++
		}

		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

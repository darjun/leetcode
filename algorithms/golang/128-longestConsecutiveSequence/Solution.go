package main

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		m[num] = struct{}{}
	}

	var maxLen int
	for _, num := range nums {
		l := 1
		n := num + 1
		for _, exist := m[n]; exist; {
			l++
			n++
			_, exist = m[n]
		}

		if l > maxLen {
			maxLen = l
		}
	}

	return maxLen
}

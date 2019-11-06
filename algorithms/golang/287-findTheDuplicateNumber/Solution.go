package main

func findDuplicate(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := m[num]; exist {
			return num
		}

		m[num] = struct{}{}
	}

	return 0
}

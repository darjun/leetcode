package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exist := m[num]; exist {
			return true
		}

		m[num] = struct{}{}
	}

	return false
}

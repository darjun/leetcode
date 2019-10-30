package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if index, exist := m[num]; exist && i-index <= k {
			return true
		} else {
			m[num] = i
		}
	}

	return false
}

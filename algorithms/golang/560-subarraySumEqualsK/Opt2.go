package main

func subarraySumOpt2(nums []int, k int) int {
	var count int

	m := make(map[int]int)

	sum := 0
	for _, num := range nums {
		sum += num

		if c, exist := m[sum - k]; exist {
			count += c
		}

		m[sum]++
	}

	return count
}
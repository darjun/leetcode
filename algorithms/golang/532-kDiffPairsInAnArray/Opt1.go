package main

func findPairsOpt1(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var total int
	for num, count := range m {
		if k == 0 {
			if count >= 2 {
				total++
			}
		} else {
			if _, exist := m[num+k]; exist {
				total++
			}
		}
	}

	return total
}
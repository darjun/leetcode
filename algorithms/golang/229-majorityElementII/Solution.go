package main

func majorityElement(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var ans []int
	for num, count := range m {
		if count > n/3 {
			ans = append(ans, num)
		}
	}
	return ans
}

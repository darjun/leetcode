package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, exist := m[target-num]; exist {
			return []int{j, i}
		}

		m[num] = i
	}

	return nil
}

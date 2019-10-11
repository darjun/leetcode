package main

import "sort"

func threeSum(nums []int) [][]int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}

	type triplet struct {
		num1, num2, num3 int
	}
	s := make(map[triplet]struct{})

	result := make([][]int, 0, 1)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			twoSum := nums[i] + nums[j]

			thirdIndex := m[-twoSum]
			if thirdIndex > j {
				t := []int{nums[i], nums[j], -twoSum}
				sort.Ints(t)

				key := triplet{t[0], t[1], t[2]}
				if _, exist := s[key]; !exist {
					result = append(result, t)
					s[key] = struct{}{}
				}
			}
		}
	}

	return result
}

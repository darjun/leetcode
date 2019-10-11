package main

import (
	"sort"
)

func threeSumOpt1(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0, 1)
	l := len(nums)
	for i := 0; i < l-2; {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				j++
				// skip same nums[j]
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				k--
				// skip same nums[k]
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			} else {
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			}
		}

		i++
		// skip same nums[i]
		for i < l-2 && nums[i] == nums[i-1] {
			i++
		}
	}

	return result
}

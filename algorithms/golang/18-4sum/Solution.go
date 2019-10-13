package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0, 1)
	l := len(nums)
	for i := 0; i < l-3; {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < l-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			p := j + 1
			q := l - 1
			for p < q {
				sum := nums[i] + nums[j] + nums[p] + nums[q]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[p], nums[q]})

					p++
					for p < q && nums[p] == nums[p-1] {
						p++
					}

					q--
					for p < q && nums[q] == nums[q+1] {
						q--
					}
				} else if sum < target {
					for p < q && nums[p] == nums[p+1] {
						p++
					}
					p++
				} else {
					for p < q && nums[q] == nums[q-1] {
						q--
					}
					q--
				}
			}
		}
	}

	return result
}

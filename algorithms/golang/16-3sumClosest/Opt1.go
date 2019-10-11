package main

import (
	"sort"
)

func threeSumClosestOpt1(nums []int, target int) int {
	sort.Ints(nums)

	minDiff := -1
	closestSum := 0
	l := len(nums)
	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			diff := sum - target
			isNegative := false
			if diff < 0 {
				diff = -diff
				isNegative = true
			}

			if diff == 0 {
				return sum
			}

			if minDiff == -1 || diff < minDiff {
				minDiff = diff
				closestSum = sum
			}

			if isNegative {
				// 负数，要向0逼近，j++
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else {
				// 正数，要向0逼近，k--
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return closestSum
}

package main

func threeSumClosest(nums []int, target int) int {
	minDiff := -1
	closestSum := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				s := nums[i] + nums[j] + nums[k]
				diff := s - target
				if diff < 0 {
					diff = -diff
				}
				if diff == 0 {
					return s
				} else if minDiff == -1 || diff < minDiff {
					minDiff = diff
					closestSum = s
				}
			}
		}
	}

	return closestSum
}

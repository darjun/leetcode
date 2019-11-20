package main

func findPairs(nums []int, k int) int {
	m := make(map[[2]int]struct{})

	for i := 0;i <len(nums)-1; i++ {
		for j := i+1;j<len(nums);j++{
			diff := nums[j] - nums[i]
			if diff < 0 {
				diff = -diff
			}

			if diff == k {
				if nums[i] > nums[j] {
					m[[2]int{nums[j], nums[i]}] = struct{}{}
				} else {
					m[[2]int{nums[i], nums[j]}] = struct{}{}
				}
			}
		}
	}

	return len(m)
}
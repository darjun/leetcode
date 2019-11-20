package main

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	sums := make([]int, len(nums), len(nums))

	var count int
	for i := 1; i < len(nums); i++ {
		if i == 0 {
            sums[i] = nums[i]
        } else {
		    sums[i] = sums[i-1] + nums[i]
        }

		if nums[i] == k {
			count++
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			var subsum int
			if i == 0 {
				subsum = sums[j]
			} else {
				subsum = sums[j] - sums[i-1]
			}

			if subsum == k {
				count++
			}
		}
	}

	return count
}